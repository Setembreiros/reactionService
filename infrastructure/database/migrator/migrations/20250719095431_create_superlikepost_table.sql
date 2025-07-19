-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE reactionservice.superlikePosts (
    postId VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL  
);

-- Índice para búsquedas por username
CREATE INDEX idx_superlikePosts_username ON reactionservice.superlikePosts(username);

-- Índice para búsquedas por postId
CREATE INDEX idx_superlikePosts_postid ON reactionservice.superlikePosts(postId);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

-- Índice para búsquedas por username
DROP INDEX idx_superlikePosts_username ON reactionservice.superlikePosts(username);

-- Índice para búsquedas por postId
DROP INDEX idx_superlikePosts_postid ON reactionservice.superlikePosts(postId);


DROP TABLE reactionservice.superlikePosts (
    postId VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL  
);
-- +goose StatementEnd
