-- +goose Up
-- +goose StatementBegin
BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL,
	uuid VARCHAR(255) NOT NULL UNIQUE,
    user_name VARCHAR(255) NOT NULL UNIQUE,
	email VARCHAR(255) NOT NULL UNIQUE,
    is_active BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    PRIMARY KEY (id)
);

CREATE INDEX idx_users_deleted_at ON users (deleted_at);

COMMIT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
BEGIN TRANSACTION;

DROP TABLE IF EXISTS users;

COMMIT;
-- +goose StatementEnd
