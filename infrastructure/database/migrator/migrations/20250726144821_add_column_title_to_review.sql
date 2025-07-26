-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE reactionservice.review
ADD COLUMN title TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE reactionservice.review
DROP COLUMN title;
-- +goose StatementEnd
