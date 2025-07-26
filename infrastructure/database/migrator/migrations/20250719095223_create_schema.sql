-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE SCHEMA reactionservice;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP SCHEMA reactionservice;
-- +goose StatementEnd
