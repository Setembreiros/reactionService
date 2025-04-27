CREATE TABLE reactionservice.superlikePosts (
    postId VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL  
);

-- Índice para búsquedas por username
CREATE INDEX idx_superlikePosts_username ON reactionservice.superlikePosts(username);

-- Índice para búsquedas por postId
CREATE INDEX idx_superlikePosts_postid ON reactionservice.superlikePosts(postId);