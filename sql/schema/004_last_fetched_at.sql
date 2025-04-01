-- +goose Up
-- +sqlc:schema
ALTER TABLE feeds
ADD COLUMN last_fetched_at TIMESTAMP NULL;