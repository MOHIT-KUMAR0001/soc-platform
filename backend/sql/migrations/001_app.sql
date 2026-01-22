-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS auth_db (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    worker_group VARCHAR(28) UNIQUE NOT NULL,
    -- Store a hash of the key, never the key itself
    api_key_hash VARCHAR(512) UNIQUE NOT NULL, 
    created_at  TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS auth_db;
-- +goose StatementEnd

