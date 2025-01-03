-- +goose Up
CREATE TABLE IF NOT EXISTS test (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS test;

-- +goose StatementBegin
-- +goose StatementEnd