package auth

import (
	"context"
	"fmt"
	"sync"

	"github.com/rilldata/rill/admin"
	"github.com/rilldata/rill/admin/pkg/authtoken"
	adminv1 "github.com/rilldata/rill/proto/gen/rill/admin/v1"
)

// OwnerType is an enum of types of claim owners
type OwnerType string

const (
	OwnerTypeAnon    OwnerType = "anon"
	OwnerTypeUser    OwnerType = "user"
	OwnerTypeService OwnerType = "service"
)

// Claims resolves permissions for a requester.
type Claims interface {
	OwnerType() OwnerType
	OwnerID() string
	AuthTokenID() string
	Superuser(ctx context.Context) bool
	OrganizationPermissions(ctx context.Context, orgID string) *adminv1.OrganizationPermissions
	ProjectPermissions(ctx context.Context, orgID, projectID string) *adminv1.ProjectPermissions
}

// claimsContextKey is used to set and get Claims on a request context.
type claimsContextKey struct{}

// GetClaims retrieves Claims from a request context.
// It should only be used in handlers intercepted by UnaryServerInterceptor or StreamServerInterceptor.
func GetClaims(ctx context.Context) Claims {
	claims, ok := ctx.Value(claimsContextKey{}).(Claims)
	if !ok {
		return nil
	}

	return claims
}

// anonClaims represents claims for an unauthenticated user.
type anonClaims struct{}

// ensure anonClaims implements Claims
var _ Claims = anonClaims{}

func (c anonClaims) OwnerType() OwnerType {
	return OwnerTypeAnon
}

func (c anonClaims) OwnerID() string {
	return ""
}

func (c anonClaims) AuthTokenID() string {
	return ""
}

func (c anonClaims) Superuser(ctx context.Context) bool {
	return false
}

func (c anonClaims) OrganizationPermissions(ctx context.Context, orgID string) *adminv1.OrganizationPermissions {
	return &adminv1.OrganizationPermissions{}
}

func (c anonClaims) ProjectPermissions(ctx context.Context, orgID, projectID string) *adminv1.ProjectPermissions {
	return &adminv1.ProjectPermissions{}
}

// authTokenClaims represents claims for an admin.AuthToken.
type authTokenClaims struct {
	sync.Mutex
	token                   admin.AuthToken
	admin                   *admin.Service
	orgPermissionsCache     map[string]*adminv1.OrganizationPermissions
	projectPermissionsCache map[string]*adminv1.ProjectPermissions
	superuser               *bool
}

func newAuthTokenClaims(token admin.AuthToken, adminService *admin.Service) Claims {
	return &authTokenClaims{
		token:                   token,
		admin:                   adminService,
		orgPermissionsCache:     make(map[string]*adminv1.OrganizationPermissions),
		projectPermissionsCache: make(map[string]*adminv1.ProjectPermissions),
		superuser:               nil,
	}
}

// ensure *authTokenClaims implements Claims
var _ Claims = &authTokenClaims{}

func (c *authTokenClaims) OwnerType() OwnerType {
	t := c.token.Token().Type
	switch t {
	case authtoken.TypeUser:
		return OwnerTypeUser
	case authtoken.TypeService:
		return OwnerTypeService
	default:
		panic(fmt.Errorf("unexpected token type %q", t))
	}
}

func (c *authTokenClaims) OwnerID() string {
	return c.token.OwnerID()
}

func (c *authTokenClaims) AuthTokenID() string {
	return c.token.Token().ID.String()
}

func (c *authTokenClaims) Superuser(ctx context.Context) bool {
	switch c.token.Token().Type {
	case authtoken.TypeUser:
		// continue
	case authtoken.TypeService:
		// services can't be superusers
		return false
	default:
		panic(fmt.Errorf("unexpected token type %q", c.token.Token().Type))
	}

	c.Lock()
	defer c.Unlock()

	if c.superuser != nil {
		return *c.superuser
	}

	user, err := c.admin.DB.FindUser(ctx, c.token.OwnerID())
	if err != nil {
		panic(fmt.Errorf("failed to get user info: %w", err))
	}

	c.superuser = &user.Superuser

	return *c.superuser
}

func (c *authTokenClaims) OrganizationPermissions(ctx context.Context, orgID string) *adminv1.OrganizationPermissions {
	c.Lock()
	defer c.Unlock()

	return c.organizationPermissionsUnsafe(ctx, orgID)
}

func (c *authTokenClaims) ProjectPermissions(ctx context.Context, orgID, projectID string) *adminv1.ProjectPermissions {
	c.Lock()
	defer c.Unlock()

	perm, ok := c.projectPermissionsCache[projectID]
	if ok {
		return perm
	}

	orgPerms := c.organizationPermissionsUnsafe(ctx, orgID)

	var err error
	switch c.token.Token().Type {
	case authtoken.TypeUser:
		perm, err = c.admin.ProjectPermissionsForUser(ctx, projectID, c.token.OwnerID(), orgPerms)
	case authtoken.TypeService:
		perm, err = c.admin.ProjectPermissionsForService(ctx, projectID, c.token.OwnerID(), orgPerms)
	default:
		err = fmt.Errorf("unexpected token type %q", c.token.Token().Type)
	}
	if err != nil {
		panic(fmt.Errorf("failed to get project permissions: %w", err))
	}

	c.projectPermissionsCache[projectID] = perm
	return perm
}

// organizationPermissionsUnsafe resolves organization permissions.
// organizationPermissionsUnsafe accesses the cache without locking, so it should only be called from a function that already has a lock.
func (c *authTokenClaims) organizationPermissionsUnsafe(ctx context.Context, orgID string) *adminv1.OrganizationPermissions {
	perm, ok := c.orgPermissionsCache[orgID]
	if ok {
		return perm
	}

	var err error
	switch c.token.Token().Type {
	case authtoken.TypeUser:
		perm, err = c.admin.OrganizationPermissionsForUser(ctx, orgID, c.token.OwnerID())
	case authtoken.TypeService:
		perm, err = c.admin.OrganizationPermissionsForService(ctx, orgID, c.token.OwnerID())
	default:
		err = fmt.Errorf("unexpected token type %q", c.token.Token().Type)
	}
	if err != nil {
		panic(fmt.Errorf("failed to get org permissions: %w", err))
	}

	c.orgPermissionsCache[orgID] = perm
	return perm
}
