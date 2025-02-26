package rillv1

import (
	"context"
	"slices"
	"strings"

	runtimev1 "github.com/rilldata/rill/proto/gen/rill/runtime/v1"
	"github.com/rilldata/rill/runtime/drivers"
	"github.com/rilldata/rill/runtime/drivers/slack"
	"go.uber.org/zap"
	"golang.org/x/exp/maps"
	"google.golang.org/protobuf/types/known/structpb"
)

// Connector contains metadata about a connector used in a Rill project
type Connector struct {
	Name            string
	Driver          string
	Spec            drivers.Spec
	DefaultConfig   map[string]string
	Resources       []*Resource
	AnonymousAccess bool
}

// AnalyzeConnectors extracts connector metadata from a Rill project
func (p *Parser) AnalyzeConnectors(ctx context.Context) ([]*Connector, error) {
	a := &connectorAnalyzer{
		parser: p,
		result: make(map[string]*Connector),
	}

	err := a.analyze(ctx)
	if err != nil {
		return nil, err
	}

	res := maps.Values(a.result)

	// Sort output to ensure deterministic ordering
	slices.SortFunc(res, func(a, b *Connector) int {
		return strings.Compare(a.Name, b.Name)
	})

	return res, nil
}

// connectorAnalyzer implements logic for extracting connector metadata from a parser
type connectorAnalyzer struct {
	parser *Parser
	result map[string]*Connector
}

// analyze is the entrypoint for connector analysis. After running it, you can access the result.
func (a *connectorAnalyzer) analyze(ctx context.Context) error {
	if a.parser.RillYAML != nil {
		// Track any connectors explicitly configured in rill.yaml
		for _, c := range a.parser.RillYAML.Connectors {
			err := a.trackConnector(c.Name, nil, false)
			if err != nil {
				return err
			}
		}

		// Track the OLAP connector specified in rill.yaml
		if a.parser.RillYAML.OLAPConnector != "" {
			err := a.trackConnector(a.parser.RillYAML.OLAPConnector, nil, false)
			if err != nil {
				return err
			}
		}
	}

	for _, r := range a.parser.Resources {
		err := a.analyzeResource(ctx, r)
		if err != nil {
			return err
		}
	}
	return nil
}

// analyzeResource extracts connector metadata for a single resource.
// NOTE: If we add more resource kinds that use connectors, add connector extraction logic here.
func (a *connectorAnalyzer) analyzeResource(ctx context.Context, r *Resource) error {
	if r.SourceSpec != nil {
		return a.analyzeSource(ctx, r)
	} else if r.ModelSpec != nil {
		return a.analyzeModel(ctx, r)
	} else if r.MetricsViewSpec != nil {
		return a.trackConnector(r.MetricsViewSpec.Connector, r, false)
	} else if r.MigrationSpec != nil {
		return a.trackConnector(r.MigrationSpec.Connector, r, false)
	} else if r.APISpec != nil {
		return a.analyzeResourceWithResolver(r, r.APISpec.Resolver, r.APISpec.ResolverProperties)
	} else if r.ComponentSpec != nil {
		return a.analyzeResourceWithResolver(r, r.ComponentSpec.Resolver, r.ComponentSpec.ResolverProperties)
	} else if r.AlertSpec != nil {
		return a.analyzeResourceNotifiers(r, r.AlertSpec.Notifiers)
	} else if r.ReportSpec != nil {
		return a.analyzeResourceNotifiers(r, r.ReportSpec.Notifiers)
	}
	// Other resource kinds currently don't use connectors.
	return nil
}

// analyzeSource extracts connector metadata for a source resource.
// The logic for extracting metadata from sources is more complex than for other resource kinds, hence the separate function.
func (a *connectorAnalyzer) analyzeSource(ctx context.Context, r *Resource) error {
	// No analysis necessary for the sink connector
	err := a.trackConnector(r.SourceSpec.SinkConnector, r, false)
	if err != nil {
		return err
	}

	// Prep for analyzing SourceConnector
	spec := r.SourceSpec
	srcProps := spec.Properties.AsMap()
	_, sourceConnector, err := a.parser.driverForConnector(spec.SourceConnector)
	if err != nil {
		return err
	}

	// Check if we have anonymous access (unless we already know that we don't)
	var anonAccess bool
	if res, ok := a.result[spec.SourceConnector]; !ok || res.AnonymousAccess {
		anonAccess, _ = sourceConnector.HasAnonymousSourceAccess(ctx, srcProps, zap.NewNop())
	}

	// Track the source connector
	err = a.trackConnector(spec.SourceConnector, r, anonAccess)
	if err != nil {
		return err
	}

	// Track any tertiary connectors (like a DuckDB source referencing S3 in its SQL).
	// NOTE: Not checking anonymous access for these since we don't know what properties to use.
	// TODO: Can we solve that issue?
	otherConnectors, _ := sourceConnector.TertiarySourceConnectors(ctx, srcProps, zap.NewNop())
	for _, connector := range otherConnectors {
		err := a.trackConnector(connector, r, false)
		if err != nil {
			return err
		}
	}

	return nil
}

// analyzeModel extracts connector metadata for a model resource.
// The logic for extracting metadata from a model is more complex than for other resource kinds, hence the separate function.
func (a *connectorAnalyzer) analyzeModel(ctx context.Context, r *Resource) error {
	// No analysis necessary for the output connector
	err := a.trackConnector(r.ModelSpec.OutputConnector, r, false)
	if err != nil {
		return err
	}

	// Prep for analyzing InputConnector
	spec := r.ModelSpec
	inputProps := spec.InputProperties.AsMap()
	_, inputConnector, err := a.parser.driverForConnector(spec.InputConnector)
	if err != nil {
		return err
	}

	// Check if we have anonymous access (unless we already know that we don't)
	var anonAccess bool
	if res, ok := a.result[spec.InputConnector]; !ok || res.AnonymousAccess {
		anonAccess, _ = inputConnector.HasAnonymousSourceAccess(ctx, inputProps, zap.NewNop())
	}

	// Track the input connector
	err = a.trackConnector(spec.InputConnector, r, anonAccess)
	if err != nil {
		return err
	}

	// Track any tertiary connectors (like a DuckDB source referencing S3 in its SQL).
	// NOTE: Not checking anonymous access for these since we don't know what properties to use.
	// TODO: Can we solve that issue?
	otherConnectors, _ := inputConnector.TertiarySourceConnectors(ctx, inputProps, zap.NewNop())
	for _, connector := range otherConnectors {
		err := a.trackConnector(connector, r, false)
		if err != nil {
			return err
		}
	}

	return nil
}

// analyzeResourceWithResolver extracts connector metadata for a resource that uses a resolver.
func (a *connectorAnalyzer) analyzeResourceWithResolver(r *Resource, resolver string, resolverProps *structpb.Struct) error {
	// The "sql" resolver takes an optional "connector" property
	if resolver == "sql" {
		for k, v := range resolverProps.Fields {
			if k == "connector" {
				connector := v.GetStringValue()
				if connector != "" {
					return a.trackConnector(connector, r, false)
				}
			}
		}
	}

	return nil
}

// analyzeResourceNotifiers extracts connector metadata for a resource that uses notifiers (email, slack, etc).
func (a *connectorAnalyzer) analyzeResourceNotifiers(r *Resource, notifiers []*runtimev1.Notifier) error {
	for _, n := range notifiers {
		anonAccess := false
		if n.Connector == "slack" {
			// Slack notifier can be used anonymously if no users and no channels are specified (only webhooks)
			props, err := slack.DecodeProps(n.Properties.AsMap())
			if err != nil {
				return err
			}
			if len(props.Users) == 0 && len(props.Channels) == 0 {
				anonAccess = true
			}
		}
		err := a.trackConnector(n.Connector, r, anonAccess)
		if err != nil {
			return err
		}
	}
	return nil
}

// trackConnector tracks a connector and an associated resource in the analyzer's result map
func (a *connectorAnalyzer) trackConnector(connector string, r *Resource, anonAccess bool) error {
	res, ok := a.result[connector]
	if !ok {
		driver, driverConnector, err := a.parser.driverForConnector(connector)
		if err != nil {
			return err
		}

		// Searfch rill.yaml for default config properties for this connector
		var defaultConfig map[string]string
		if a.parser.RillYAML != nil {
			for _, c := range a.parser.RillYAML.Connectors {
				if c.Name == connector {
					defaultConfig = c.Defaults
					break
				}
			}
		}

		res = &Connector{
			Name:            connector,
			Driver:          driver,
			Spec:            driverConnector.Spec(),
			DefaultConfig:   defaultConfig,
			AnonymousAccess: true,
		}

		a.result[connector] = res
	}

	if r != nil {
		found := false
		for _, existing := range res.Resources {
			if r.Name.Normalized() == existing.Name.Normalized() {
				found = true
				break
			}
		}
		if !found {
			res.Resources = append(res.Resources, r)
		}
	}

	if !anonAccess {
		res.AnonymousAccess = false
	}

	return nil
}
