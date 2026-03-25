-- name: CreateOrganization :one
INSERT INTO organizations (name, slug, status)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetOrganizationByID :one
SELECT * FROM organizations
WHERE id = $1;

-- name: GetOrganizationBySlug :one
SELECT * FROM organizations
WHERE slug = $1;

-- name: ListOrganizations :many
SELECT * FROM organizations
ORDER BY name ASC
LIMIT $1
OFFSET $2;

-- name: ListOrganizationsByStatus :many
SELECT * FROM organizations
WHERE status = $1
ORDER BY name ASC
LIMIT $2
OFFSET $3;

-- name: CountOrganizations :one
SELECT COUNT(*) FROM organizations;

-- name: CountOrganizationsByStatus :one
SELECT COUNT(*) FROM organizations
WHERE status = $1;

-- name: UpdateOrganization :one
UPDATE organizations
SET
    name       = $2,
    slug       = $3,
    status     = $4,
    updated_at = NOW()
WHERE id = $1
RETURNING *;
