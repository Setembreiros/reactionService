package unit_test_unlike_post

import (
	"errors"
	"testing"

	database "reactionservice/internal/db"
	mock_database "reactionservice/internal/db/test/mock"
	"reactionservice/internal/feature/unlike_post"
	model "reactionservice/internal/model/domain"

	"github.com/stretchr/testify/assert"
)

var dbClient *mock_database.MockDatabaseClient
var deleteLikePostRepository *unlike_post.DeleteLikePostRepository

func repositorySetUp(t *testing.T) {
	setUp(t)
	dbClient = mock_database.NewMockDatabaseClient(ctrl)
	deleteLikePostRepository = unlike_post.NewDeleteLikePostRepository(database.NewDatabase(dbClient))
}

func TestDeleteLikePostInRepository_WhenItReturnsSuccess(t *testing.T) {
	repositorySetUp(t)
	like := &model.LikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	dbClient.EXPECT().DeleteLikePost(like).Return(nil)

	err := deleteLikePostRepository.DeleteLikePost(like)

	assert.Nil(t, err)
}

func TestErrorOnDeleteLikePostInRepository_WhenDeleteLikePostFails(t *testing.T) {
	repositorySetUp(t)
	like := &model.LikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	dbClient.EXPECT().DeleteLikePost(like).Return(errors.New("some error"))

	err := deleteLikePostRepository.DeleteLikePost(like)

	assert.NotNil(t, err)
}
