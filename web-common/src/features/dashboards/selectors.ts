import { filterExpressions } from "@rilldata/web-common/features/dashboards/stores/filter-utils";
import {
  ResourceKind,
  useFilteredResources,
  useResource,
} from "@rilldata/web-common/features/entity-management/resource-selectors";
import {
  V1Expression,
  V1MetricsViewSpec,
  createQueryServiceMetricsViewTimeRange,
  createRuntimeServiceListResources,
  type V1MetricsViewTimeRangeResponse,
} from "@rilldata/web-common/runtime-client";
import type {
  CreateQueryOptions,
  CreateQueryResult,
} from "@tanstack/svelte-query";
import { derived } from "svelte/store";

export function useDashboard(instanceId: string, metricViewName: string) {
  return useResource(instanceId, metricViewName, ResourceKind.MetricsView);
}

export function useValidDashboards(instanceId: string) {
  return useFilteredResources(instanceId, ResourceKind.MetricsView, (data) =>
    data?.resources?.filter((res) => !!res.metricsView?.state?.validSpec),
  );
}

export function useValidVisualizations(instanceId: string) {
  return createRuntimeServiceListResources(
    instanceId,
    undefined, // TODO: it'd be nice if we could provide multiple kinds here
    {
      query: {
        select: (data) => {
          // Filter for valid Metrics Explorers and all Custom Dashboards (which don't yet have a valid/invalid state)
          return data?.resources?.filter(
            (res) => !!res.metricsView?.state?.validSpec || res.dashboard,
          );
        },
      },
    },
  );
}

/**
 * Gets the valid metrics view spec. Only to be used in displaying a dashboard.
 * Use {@link useDashboard} in the metrics view editor and other use cases.
 */
export const useMetricsView = <T = V1MetricsViewSpec>(
  instanceId: string,
  metricViewName: string,
  selector?: (meta: V1MetricsViewSpec) => T,
) => {
  return useResource<T>(
    instanceId,
    metricViewName,
    ResourceKind.MetricsView,
    (data) =>
      selector
        ? selector(data.metricsView?.state?.validSpec)
        : (data.metricsView?.state?.validSpec as T),
  );
};

// TODO: cleanup usage of useModelHasTimeSeries and useModelAllTimeRange
export const useModelHasTimeSeries = (
  instanceId: string,
  metricViewName: string,
) =>
  useMetricsView(instanceId, metricViewName, (meta) => !!meta?.timeDimension);

export function useMetricsViewTimeRange(
  instanceId: string,
  metricsViewName: string,
  options?: {
    query?: CreateQueryOptions<V1MetricsViewTimeRangeResponse>;
  },
): CreateQueryResult<V1MetricsViewTimeRangeResponse> {
  const { query: queryOptions } = options ?? {};

  return derived(
    [useMetricsView(instanceId, metricsViewName)],
    ([metricsView], set) =>
      createQueryServiceMetricsViewTimeRange(
        instanceId,
        metricsViewName,
        {},
        {
          query: {
            ...queryOptions,
            enabled: !!metricsView.data?.timeDimension && queryOptions?.enabled,
          },
        },
      ).subscribe(set),
  );
}

export const useMetaMeasure = (
  instanceId: string,
  metricViewName: string,
  measureName: string,
) =>
  useMetricsView(instanceId, metricViewName, (meta) =>
    meta?.measures?.find((measure) => measure.name === measureName),
  );

export const useMetaDimension = (
  instanceId: string,
  metricViewName: string,
  dimensionName: string,
) =>
  useMetricsView(instanceId, metricViewName, (meta) => {
    const dim = meta?.dimensions?.find(
      (dimension) => dimension.name === dimensionName,
    );
    return {
      ...dim,
      // this is for backwards compatibility when we used `name` as `column`
      column: dim.column ?? dim.name,
    };
  });

/**
 * Returns a copy of a V1MetricsViewFilter that does not include
 * the filters for the specified dimension name.
 */
export const getFiltersForOtherDimensions = (
  filters: V1Expression,
  dimensionName: string,
) => {
  if (!filters) return undefined;
  return filterExpressions(
    filters,
    (e) => e.cond?.exprs?.[0].ident !== dimensionName,
  );
};

export const useGetDashboardsForModel = (
  instanceId: string,
  modelName: string,
) => {
  return useFilteredResources(instanceId, ResourceKind.MetricsView, (data) =>
    data.resources.filter((res) => res.metricsView?.spec?.table === modelName),
  );
};
