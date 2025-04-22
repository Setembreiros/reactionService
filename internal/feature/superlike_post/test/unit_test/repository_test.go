package unit_test_superlike_post

import (
	"errors"
	"testing"

	database "reactionservice/internal/db"
	mock_database "reactionservice/internal/db/test/mock"
	"reactionservice/internal/feature/superlike_post"
	model "reactionservice/internal/model/domain"

	"github.com/stretchr/testify/assert"
)

var dbClient *mock_database.MockDatabaseClient
var createSuperlikePostRepository *superlike_post.CreateSuperlikePostRepository

func repositorySetUp(t *testing.T) {
	setUp(t)
	dbClient = mock_database.NewMockDatabaseClient(ctrl)
	createSuperlikePostRepository = superlike_post.NewCreateSuperlikePostRepository(database.NewDatabase(dbClient))
}

func TestCreateSuperlikePostInRepository_WhenItReturnsSuccess(t *testing.T) {
	repositorySetUp(t)
	superlike := &model.SuperlikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	dbClient.EXPECT().GetSuperlikePost(superlike.PostId, superlike.Username).Return(nil, nil)
	dbClient.EXPECT().CreateSuperlikePost(superlike).Return(nil)

	err := createSuperlikePostRepository.CreateSuperlikePost(superlike)

	assert.Nil(t, err)
}

func TestErrorOnCreateSuperlikePostInRepository_WhenCreateSuperlikePostFails(t *testing.T) {
	repositorySetUp(t)
	superlike := &model.SuperlikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	dbClient.EXPECT().GetSuperlikePost(superlike.PostId, superlike.Username).Return(nil, nil)
	dbClient.EXPECT().CreateSuperlikePost(superlike).Return(errors.New("some error"))

	err := createSuperlikePostRepository.CreateSuperlikePost(superlike)

	assert.NotNil(t, err)
}
