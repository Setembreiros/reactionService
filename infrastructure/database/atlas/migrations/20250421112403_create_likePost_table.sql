CREATE TABLE reactionservice.likePosts (
    postId VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL  
);

-- Índice para búsquedas por username
CREATE INDEX idx_likePosts_username ON reactionservice.likePosts(username);

-- Índice para búsquedas por postId
CREATE INDEX idx_likePosts_postid ON reactionservice.likePosts(postId);