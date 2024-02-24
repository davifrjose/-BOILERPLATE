-- +goose Up
ALTER TABLE IF EXISTS users
ADD COLUMN IF NOT EXISTS api_key VARCHAR(255);
-- +goose Down
ALTER TABLE users
DROP COLUMN api_key;

