package fixtures

import (
	"context"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/google/uuid"
)

func NewContext() context.Context {
	return context.Background()
}

func NewContextWithTenant(tenantID domain.TenantID) context.Context {
	ctx := NewContext()
	return domain.TenantIDToContext(ctx, tenantID)
}

func NewContextWithUserClaims(claims domain.UserClaims) context.Context {
	ctx := NewContext()
	ctx = domain.ClaimsToContext(ctx, claims)
	ctx = domain.TenantIDToContext(ctx, claims.TenantID)
	return ctx
}

func NewTestTenantID() domain.TenantID {
	return domain.TenantID(uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"))
}

func NewTestUserClaims() domain.UserClaims {
	return domain.UserClaims{
		UserID:   domain.UserID(uuid.MustParse("660e8400-e29b-41d4-a716-446655440000")),
		TenantID: NewTestTenantID(),
		Role:     "admin",
	}
}

func NewContextWithAdminUser() context.Context {
	return NewContextWithUserClaims(NewTestUserClaims())
}
