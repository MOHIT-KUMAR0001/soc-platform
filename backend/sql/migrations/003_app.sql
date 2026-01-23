-- +goose Up
-- +goose StatementBegin
-- 1. Add the functional columns first
ALTER TABLE auth_db 
ADD COLUMN IF NOT EXISTS is_disabled BOOLEAN DEFAULT FALSE,
ADD COLUMN IF NOT EXISTS expiry TIMESTAMPTZ;

-- 2. Add updated_at at the very end
ALTER TABLE auth_db 
ADD COLUMN IF NOT EXISTS updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE auth_db 
DROP COLUMN IF EXISTS updated_at,
DROP COLUMN IF EXISTS expiry,
DROP COLUMN IF EXISTS is_disabled;
-- +goose StatementEnd