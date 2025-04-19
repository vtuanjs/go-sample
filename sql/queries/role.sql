-- name: CreateRole :one
INSERT INTO roles (name, description, note) VALUES ($1, $2, $3) RETURNING *;

-- name: GetRole :one
SELECT id, name, description, note, created_at, updated_at FROM roles WHERE id = $1 AND deleted_at IS NULL LIMIT 1;

-- name: GetRoles :many
SELECT id, name, description, note, created_at, updated_at FROM roles WHERE deleted_at IS NULL;

-- name: UpdateRole :one
UPDATE roles SET name = $2, description = $3, note = $4 WHERE id = $1 RETURNING *;

-- name: DeleteRole :one
DELETE FROM roles WHERE id = $1 RETURNING *;
