-- +goose Up
-- +goose StatementBegin
ALTER TABLE tasks ADD COLUMN image_url TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tasks DROP COLUMN image_url;
-- +goose StatementEnd 