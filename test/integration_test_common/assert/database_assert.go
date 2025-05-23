package integration_test_assert

import (
	database "reactionservice/internal/db"
	model "reactionservice/internal/model/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertLikePostExists(t *testing.T, db *database.Database, expectedLikePost *model.LikePost) {
	like, err := db.Client.GetLikePost(expectedLikePost.PostId, expectedLikePost.Username)
	assert.Nil(t, err)
	assert.Equal(t, expectedLikePost.PostId, like.PostId)
	assert.Equal(t, expectedLikePost.Username, like.Username)
}

func AssertLikePostDoesNotExists(t *testing.T, db *database.Database, like *model.LikePost) {
	like, err := db.Client.GetLikePost(like.PostId, like.Username)
	assert.Nil(t, err)
	assert.Nil(t, like)
}

func AssertSuperlikePostExists(t *testing.T, db *database.Database, expectedSuperlikePost *model.SuperlikePost) {
	like, err := db.Client.GetSuperlikePost(expectedSuperlikePost.PostId, expectedSuperlikePost.Username)
	assert.Nil(t, err)
	assert.Equal(t, expectedSuperlikePost.PostId, like.PostId)
	assert.Equal(t, expectedSuperlikePost.Username, like.Username)
}

func AssertSuperlikePostDoesNotExists(t *testing.T, db *database.Database, like *model.SuperlikePost) {
	like, err := db.Client.GetSuperlikePost(like.PostId, like.Username)
	assert.Nil(t, err)
	assert.Nil(t, like)
}
