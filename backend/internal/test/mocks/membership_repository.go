package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	"github.com/dsbraz/bud2/backend/internal/domain/membership"
)

type MembershipRepository struct {
	mock.Mock
}

func (m *MembershipRepository) Create(ctx context.Context, member *membership.Membership) (*membership.Membership, error) {
	args := m.Called(ctx, member)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*membership.Membership), args.Error(1)
}

func (m *MembershipRepository) GetByID(ctx context.Context, id uuid.UUID) (*membership.Membership, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*membership.Membership), args.Error(1)
}

func (m *MembershipRepository) GetByOrganizationAndUser(ctx context.Context, organizationID, userID uuid.UUID) (*membership.Membership, error) {
	args := m.Called(ctx, organizationID, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*membership.Membership), args.Error(1)
}

func (m *MembershipRepository) ListByOrganization(ctx context.Context, filter membership.ListByOrganizationFilter) (membership.ListResult, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(membership.ListResult), args.Error(1)
}

func (m *MembershipRepository) ListByUser(ctx context.Context, filter membership.ListByUserFilter) (membership.ListResult, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(membership.ListResult), args.Error(1)
}

func (m *MembershipRepository) Update(ctx context.Context, member *membership.Membership) (*membership.Membership, error) {
	args := m.Called(ctx, member)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*membership.Membership), args.Error(1)
}
