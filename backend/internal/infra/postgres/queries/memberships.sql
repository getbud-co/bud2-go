-- name: CreateOrganizationMembership :one
INSERT INTO organization_memberships (organization_id, user_id, role, status)
VALUES ($1, $2, $3, $4)
RETURNING id, organization_id, user_id, role, status, invited_by_user_id, joined_at, created_at, updated_at;

-- name: GetOrganizationMembershipByID :one
SELECT id, organization_id, user_id, role, status, invited_by_user_id, joined_at, created_at, updated_at
FROM organization_memberships
WHERE id = $1
  AND deleted_at IS NULL;

-- name: GetOrganizationMembershipByOrganizationAndUser :one
SELECT id, organization_id, user_id, role, status, invited_by_user_id, joined_at, created_at, updated_at
FROM organization_memberships
WHERE organization_id = $1
  AND user_id = $2
  AND deleted_at IS NULL;

-- name: ListOrganizationMemberships :many
SELECT id, organization_id, user_id, role, status, invited_by_user_id, joined_at, created_at, updated_at
FROM organization_memberships
WHERE organization_id = $1
  AND deleted_at IS NULL
ORDER BY created_at ASC
LIMIT $2 OFFSET $3;

-- name: ListOrganizationMembershipsByStatus :many
SELECT id, organization_id, user_id, role, status, invited_by_user_id, joined_at, created_at, updated_at
FROM organization_memberships
WHERE organization_id = $1
  AND status = $2
  AND deleted_at IS NULL
ORDER BY created_at ASC
LIMIT $3 OFFSET $4;

-- name: CountOrganizationMemberships :one
SELECT COUNT(*) FROM organization_memberships
WHERE organization_id = $1
  AND deleted_at IS NULL;

-- name: CountOrganizationMembershipsByStatus :one
SELECT COUNT(*) FROM organization_memberships
WHERE organization_id = $1
  AND status = $2
  AND deleted_at IS NULL;

-- name: ListUserMemberships :many
SELECT id, organization_id, user_id, role, status, invited_by_user_id, joined_at, created_at, updated_at
FROM organization_memberships
WHERE user_id = $1
  AND deleted_at IS NULL
ORDER BY created_at ASC
LIMIT $2 OFFSET $3;

-- name: ListUserMembershipsByStatus :many
SELECT id, organization_id, user_id, role, status, invited_by_user_id, joined_at, created_at, updated_at
FROM organization_memberships
WHERE user_id = $1
  AND status = $2
  AND deleted_at IS NULL
ORDER BY created_at ASC
LIMIT $3 OFFSET $4;

-- name: CountUserMemberships :one
SELECT COUNT(*) FROM organization_memberships
WHERE user_id = $1
  AND deleted_at IS NULL;

-- name: CountUserMembershipsByStatus :one
SELECT COUNT(*) FROM organization_memberships
WHERE user_id = $1
  AND status = $2
  AND deleted_at IS NULL;

-- name: UpdateOrganizationMembership :one
UPDATE organization_memberships
SET role = $2,
    status = $3,
    updated_at = NOW()
WHERE id = $1
  AND deleted_at IS NULL
RETURNING id, organization_id, user_id, role, status, invited_by_user_id, joined_at, created_at, updated_at;

-- name: SoftDeleteOrganizationMembership :exec
UPDATE organization_memberships
SET deleted_at = NOW()
WHERE id = $1
  AND deleted_at IS NULL;
