-- name: CreateUser :one
INSERT INTO users (id, tenant_id, name, email, role, status)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 AND tenant_id = $2;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 AND tenant_id = $2;

-- name: ListUsers :many
SELECT * FROM users
WHERE tenant_id = $1
ORDER BY name ASC
LIMIT $2 OFFSET $3;

-- name: ListUsersByStatus :many
SELECT * FROM users
WHERE tenant_id = $1
  AND status = $2
ORDER BY name ASC
LIMIT $3 OFFSET $4;

-- name: SearchUsers :many
SELECT * FROM users
WHERE tenant_id = $1
  AND (name ILIKE $2 OR email ILIKE $2)
ORDER BY name ASC
LIMIT $3 OFFSET $4;

-- name: CountUsers :one
SELECT COUNT(*) FROM users
WHERE tenant_id = $1;

-- name: CountUsersByStatus :one
SELECT COUNT(*) FROM users
WHERE tenant_id = $1 AND status = $2;

-- name: CountSearchUsers :one
SELECT COUNT(*) FROM users
WHERE tenant_id = $1
  AND (name ILIKE $2 OR email ILIKE $2);

-- name: UpdateUser :one
UPDATE users
SET name       = $3,
    email      = $4,
    role       = $5,
    status     = $6,
    updated_at = NOW()
WHERE id = $1 AND tenant_id = $2
RETURNING *;
