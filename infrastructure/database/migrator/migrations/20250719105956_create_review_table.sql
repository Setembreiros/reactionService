-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE reactionservice.review (
    id BIGSERIAL PRIMARY KEY,
    postId VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    rating INTEGER NOT NULL CHECK (rating >= 0 AND rating <= 5),
    createdAt TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Índice para búsquedas por username
CREATE INDEX idx_review_username ON reactionservice.review(username);

-- Índice para búsquedas por postId
CREATE INDEX idx_review_postid ON reactionservice.review(postId);

-- Índice compuesto para búsquedas que combinan username y postId
CREATE INDEX idx_review_username_postid ON reactionservice.review(username, postId);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP INDEX idx_review_username;
DROP INDEX idx_review_postid;
DROP INDEX idx_review_username_postid;

DROP TABLE reactionservice.review;
-- +goose StatementEnd
