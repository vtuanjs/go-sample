-- +goose Up
-- +goose StatementBegin
BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS roles (
    id BIGSERIAL,
    name VARCHAR(255) NOT NULL UNIQUE,
    description VARCHAR(255) NOT NULL,
    note VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    PRIMARY KEY (id)
);

CREATE INDEX idx_roles_deleted_at ON roles (deleted_at);

COMMIT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS roles;
-- +goose StatementEnd
