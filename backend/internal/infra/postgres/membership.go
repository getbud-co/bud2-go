package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/getbud-co/bud2/backend/internal/domain/membership"
	"github.com/getbud-co/bud2/backend/internal/infra/postgres/sqlc"
)

type membershipQuerier interface {
	CreateOrganizationMembership(ctx context.Context, arg sqlc.CreateOrganizationMembershipParams) (sqlc.CreateOrganizationMembershipRow, error)
	GetOrganizationMembershipByID(ctx context.Context, id uuid.UUID) (sqlc.GetOrganizationMembershipByIDRow, error)
	GetOrganizationMembershipByOrganizationAndUser(ctx context.Context, arg sqlc.GetOrganizationMembershipByOrganizationAndUserParams) (sqlc.GetOrganizationMembershipByOrganizationAndUserRow, error)
	ListOrganizationMemberships(ctx context.Context, arg sqlc.ListOrganizationMembershipsParams) ([]sqlc.ListOrganizationMembershipsRow, error)
	ListOrganizationMembershipsByStatus(ctx context.Context, arg sqlc.ListOrganizationMembershipsByStatusParams) ([]sqlc.ListOrganizationMembershipsByStatusRow, error)
	CountOrganizationMemberships(ctx context.Context, organizationID uuid.UUID) (int64, error)
	CountOrganizationMembershipsByStatus(ctx context.Context, arg sqlc.CountOrganizationMembershipsByStatusParams) (int64, error)
	ListUserMemberships(ctx context.Context, arg sqlc.ListUserMembershipsParams) ([]sqlc.ListUserMembershipsRow, error)
	ListUserMembershipsByStatus(ctx context.Context, arg sqlc.ListUserMembershipsByStatusParams) ([]sqlc.ListUserMembershipsByStatusRow, error)
	CountUserMemberships(ctx context.Context, userID uuid.UUID) (int64, error)
	CountUserMembershipsByStatus(ctx context.Context, arg sqlc.CountUserMembershipsByStatusParams) (int64, error)
	UpdateOrganizationMembership(ctx context.Context, arg sqlc.UpdateOrganizationMembershipParams) (sqlc.UpdateOrganizationMembershipRow, error)
}

type MembershipRepository struct {
	q membershipQuerier
}

func NewMembershipRepository(q membershipQuerier) *MembershipRepository {
	return &MembershipRepository{q: q}
}

func (r *MembershipRepository) Create(ctx context.Context, m *membership.Membership) (*membership.Membership, error) {
	row, err := r.q.CreateOrganizationMembership(ctx, sqlc.CreateOrganizationMembershipParams{
		OrganizationID: m.OrganizationID,
		UserID:         m.UserID,
		Role:           string(m.Role),
		Status:         string(m.Status),
	})
	if err != nil {
		return nil, err
	}
	return createMembershipRowToDomain(row), nil
}

func (r *MembershipRepository) GetByID(ctx context.Context, id uuid.UUID) (*membership.Membership, error) {
	row, err := r.q.GetOrganizationMembershipByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, membership.ErrNotFound
		}
		return nil, err
	}
	return getMembershipByIDRowToDomain(row), nil
}

func (r *MembershipRepository) GetByOrganizationAndUser(ctx context.Context, organizationID, userID uuid.UUID) (*membership.Membership, error) {
	row, err := r.q.GetOrganizationMembershipByOrganizationAndUser(ctx, sqlc.GetOrganizationMembershipByOrganizationAndUserParams{
		OrganizationID: organizationID,
		UserID:         userID,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, membership.ErrNotFound
		}
		return nil, err
	}
	return getMembershipByOrganizationAndUserRowToDomain(row), nil
}

func (r *MembershipRepository) ListByOrganization(ctx context.Context, filter membership.ListByOrganizationFilter) (membership.ListResult, error) {
	limit := int32(filter.Size)
	offset := int32((filter.Page - 1) * filter.Size)

	var items []membership.Membership
	var total int64
	var err error

	if filter.Status != nil {
		rows, listErr := r.q.ListOrganizationMembershipsByStatus(ctx, sqlc.ListOrganizationMembershipsByStatusParams{
			OrganizationID: filter.OrganizationID,
			Status:         string(*filter.Status),
			Limit:          limit,
			Offset:         offset,
		})
		if listErr != nil {
			return membership.ListResult{}, listErr
		}
		items = make([]membership.Membership, len(rows))
		for i, row := range rows {
			items[i] = *listOrganizationMembershipsByStatusRowToDomain(row)
		}
		total, err = r.q.CountOrganizationMembershipsByStatus(ctx, sqlc.CountOrganizationMembershipsByStatusParams{
			OrganizationID: filter.OrganizationID,
			Status:         string(*filter.Status),
		})
	} else {
		rows, listErr := r.q.ListOrganizationMemberships(ctx, sqlc.ListOrganizationMembershipsParams{
			OrganizationID: filter.OrganizationID,
			Limit:          limit,
			Offset:         offset,
		})
		if listErr != nil {
			return membership.ListResult{}, listErr
		}
		items = make([]membership.Membership, len(rows))
		for i, row := range rows {
			items[i] = *listOrganizationMembershipsRowToDomain(row)
		}
		total, err = r.q.CountOrganizationMemberships(ctx, filter.OrganizationID)
	}
	if err != nil {
		return membership.ListResult{}, err
	}
	return membership.ListResult{Memberships: items, Total: total}, nil
}

func (r *MembershipRepository) ListByUser(ctx context.Context, filter membership.ListByUserFilter) (membership.ListResult, error) {
	limit := int32(filter.Size)
	offset := int32((filter.Page - 1) * filter.Size)

	var items []membership.Membership
	var total int64
	var err error

	if filter.Status != nil {
		rows, listErr := r.q.ListUserMembershipsByStatus(ctx, sqlc.ListUserMembershipsByStatusParams{
			UserID: filter.UserID,
			Status: string(*filter.Status),
			Limit:  limit,
			Offset: offset,
		})
		if listErr != nil {
			return membership.ListResult{}, listErr
		}
		items = make([]membership.Membership, len(rows))
		for i, row := range rows {
			items[i] = *listUserMembershipsByStatusRowToDomain(row)
		}
		total, err = r.q.CountUserMembershipsByStatus(ctx, sqlc.CountUserMembershipsByStatusParams{UserID: filter.UserID, Status: string(*filter.Status)})
	} else {
		rows, listErr := r.q.ListUserMemberships(ctx, sqlc.ListUserMembershipsParams{UserID: filter.UserID, Limit: limit, Offset: offset})
		if listErr != nil {
			return membership.ListResult{}, listErr
		}
		items = make([]membership.Membership, len(rows))
		for i, row := range rows {
			items[i] = *listUserMembershipsRowToDomain(row)
		}
		total, err = r.q.CountUserMemberships(ctx, filter.UserID)
	}
	if err != nil {
		return membership.ListResult{}, err
	}
	return membership.ListResult{Memberships: items, Total: total}, nil
}

func (r *MembershipRepository) Update(ctx context.Context, m *membership.Membership) (*membership.Membership, error) {
	row, err := r.q.UpdateOrganizationMembership(ctx, sqlc.UpdateOrganizationMembershipParams{ID: m.ID, Role: string(m.Role), Status: string(m.Status)})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, membership.ErrNotFound
		}
		return nil, err
	}
	return updateMembershipRowToDomain(row), nil
}

func createMembershipRowToDomain(row sqlc.CreateOrganizationMembershipRow) *membership.Membership {
	return membershipRowToDomain(row.ID, row.OrganizationID, row.UserID, row.Role, row.Status, row.InvitedByUserID, row.JoinedAt, row.CreatedAt, row.UpdatedAt)
}

func getMembershipByIDRowToDomain(row sqlc.GetOrganizationMembershipByIDRow) *membership.Membership {
	return membershipRowToDomain(row.ID, row.OrganizationID, row.UserID, row.Role, row.Status, row.InvitedByUserID, row.JoinedAt, row.CreatedAt, row.UpdatedAt)
}

func getMembershipByOrganizationAndUserRowToDomain(row sqlc.GetOrganizationMembershipByOrganizationAndUserRow) *membership.Membership {
	return membershipRowToDomain(row.ID, row.OrganizationID, row.UserID, row.Role, row.Status, row.InvitedByUserID, row.JoinedAt, row.CreatedAt, row.UpdatedAt)
}

func listOrganizationMembershipsRowToDomain(row sqlc.ListOrganizationMembershipsRow) *membership.Membership {
	return membershipRowToDomain(row.ID, row.OrganizationID, row.UserID, row.Role, row.Status, row.InvitedByUserID, row.JoinedAt, row.CreatedAt, row.UpdatedAt)
}

func listOrganizationMembershipsByStatusRowToDomain(row sqlc.ListOrganizationMembershipsByStatusRow) *membership.Membership {
	return membershipRowToDomain(row.ID, row.OrganizationID, row.UserID, row.Role, row.Status, row.InvitedByUserID, row.JoinedAt, row.CreatedAt, row.UpdatedAt)
}

func listUserMembershipsRowToDomain(row sqlc.ListUserMembershipsRow) *membership.Membership {
	return membershipRowToDomain(row.ID, row.OrganizationID, row.UserID, row.Role, row.Status, row.InvitedByUserID, row.JoinedAt, row.CreatedAt, row.UpdatedAt)
}

func listUserMembershipsByStatusRowToDomain(row sqlc.ListUserMembershipsByStatusRow) *membership.Membership {
	return membershipRowToDomain(row.ID, row.OrganizationID, row.UserID, row.Role, row.Status, row.InvitedByUserID, row.JoinedAt, row.CreatedAt, row.UpdatedAt)
}

func updateMembershipRowToDomain(row sqlc.UpdateOrganizationMembershipRow) *membership.Membership {
	return membershipRowToDomain(row.ID, row.OrganizationID, row.UserID, row.Role, row.Status, row.InvitedByUserID, row.JoinedAt, row.CreatedAt, row.UpdatedAt)
}

func membershipRowToDomain(id, organizationID, userID uuid.UUID, role, status string, invitedByUserID pgtype.UUID, joinedAt pgtype.Timestamptz, createdAt, updatedAt time.Time) *membership.Membership {
	result := &membership.Membership{
		ID:             id,
		OrganizationID: organizationID,
		UserID:         userID,
		Role:           membership.Role(role),
		Status:         membership.Status(status),
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}
	if invitedByUserID.Valid {
		value := uuid.UUID(invitedByUserID.Bytes)
		result.InvitedByUserID = &value
	}
	if joinedAt.Valid {
		value := joinedAt.Time
		result.JoinedAt = &value
	}
	return result
}
