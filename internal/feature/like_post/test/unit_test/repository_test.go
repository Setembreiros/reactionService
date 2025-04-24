package unit_test_like_post

import (
	"errors"
	"testing"

	database "reactionservice/internal/db"
	mock_database "reactionservice/internal/db/test/mock"
	"reactionservice/internal/feature/like_post"
	model "reactionservice/internal/model/domain"

	"github.com/stretchr/testify/assert"
)

var dbClient *mock_database.MockDatabaseClient
var createLikePostRepository *like_post.CreateLikePostRepository

func repositorySetUp(t *testing.T) {
	setUp(t)
	dbClient = mock_database.NewMockDatabaseClient(ctrl)
	createLikePostRepository = like_post.NewCreateLikePostRepository(database.NewDatabase(dbClient))
}

func TestCreateLikePostInRepository_WhenItReturnsSuccess(t *testing.T) {
	repositorySetUp(t)
	like := &model.LikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	dbClient.EXPECT().GetLikePost(like.PostId, like.Username).Return(nil, nil)
	dbClient.EXPECT().CreateLikePost(like).Return(nil)

	err := createLikePostRepository.CreateLikePost(like)

	assert.Nil(t, err)
}

func TestErrorOnCreateLikePostInRepository_WhenCreateLikePostFails(t *testing.T) {
	repositorySetUp(t)
	like := &model.LikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	dbClient.EXPECT().GetLikePost(like.PostId, like.Username).Return(nil, nil)
	dbClient.EXPECT().CreateLikePost(like).Return(errors.New("some error"))

	err := createLikePostRepository.CreateLikePost(like)

	assert.NotNil(t, err)
}
