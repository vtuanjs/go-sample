-- name: CreateUser :one
INSERT INTO users (uuid, user_name, email, is_active) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetUser :one
SELECT id, uuid, user_name, email, is_active, created_at, updated_at FROM users WHERE id = $1 AND deleted_at IS NULL LIMIT 1;

-- name: GetUsers :many
SELECT id, uuid, user_name, email, is_active, created_at, updated_at FROM users WHERE deleted_at IS NULL;

-- name: UpdateUser :one
UPDATE users SET user_name = $2, is_active = $3 WHERE id = $1 RETURNING *;

-- name: DeleteUser :one
DELETE FROM users WHERE id = $1 RETURNING *;

-- name: GetUserByEmail :one
SELECT id, uuid, user_name, email, is_active, created_at, updated_at FROM users WHERE user_name = $1 AND deleted_at IS NULL LIMIT 1;
