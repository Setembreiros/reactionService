package unit_test_unsuperlike_post

import (
	"errors"
	"testing"

	database "reactionservice/internal/db"
	mock_database "reactionservice/internal/db/test/mock"
	"reactionservice/internal/feature/unsuperlike_post"
	model "reactionservice/internal/model/domain"

	"github.com/stretchr/testify/assert"
)

var dbClient *mock_database.MockDatabaseClient
var deleteSuperlikePostRepository *unsuperlike_post.DeleteSuperlikePostRepository

func repositorySetUp(t *testing.T) {
	setUp(t)
	dbClient = mock_database.NewMockDatabaseClient(ctrl)
	deleteSuperlikePostRepository = unsuperlike_post.NewDeleteSuperlikePostRepository(database.NewDatabase(dbClient))
}

func TestDeleteSuperlikePostInRepository_WhenItReturnsSuccess(t *testing.T) {
	repositorySetUp(t)
	superlike := &model.SuperlikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	dbClient.EXPECT().DeleteSuperlikePost(superlike).Return(nil)

	err := deleteSuperlikePostRepository.DeleteSuperlikePost(superlike)

	assert.Nil(t, err)
}

func TestErrorOnDeleteSuperlikePostInRepository_WhenDeleteSuperlikePostFails(t *testing.T) {
	repositorySetUp(t)
	superlike := &model.SuperlikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	dbClient.EXPECT().DeleteSuperlikePost(superlike).Return(errors.New("some error"))

	err := deleteSuperlikePostRepository.DeleteSuperlikePost(superlike)

	assert.NotNil(t, err)
}
