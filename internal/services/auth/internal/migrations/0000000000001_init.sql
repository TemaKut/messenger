-- +goose Up
CREATE TABLE users (
    id        UUID,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    phone VARCHAR(30),
    props JSONB,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    PRIMARY KEY (id)
);
-- +goose Down
DROP TABLE IF EXISTS users;