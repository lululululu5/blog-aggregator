-- +goose Up
ALTER TABLE feed_follows
RENAME COLUMN crated_at to created_at;

-- +goose Down
ALTER TABLE feed_follows
RENAME COLUMN created_at to crated_at;