package unit_test_unsuperlike_post

import (
	"errors"
	"fmt"
	"testing"

	"reactionservice/internal/bus"
	mock_bus "reactionservice/internal/bus/test/mock"
	"reactionservice/internal/feature/unsuperlike_post"
	mock_unsuperlike_post "reactionservice/internal/feature/unsuperlike_post/test/mock"
	model "reactionservice/internal/model/domain"
	"reactionservice/internal/model/event"

	"github.com/stretchr/testify/assert"
)

var serviceRepository *mock_unsuperlike_post.MockRepository
var serviceExternalBus *mock_bus.MockExternalBus
var serviceBus *bus.EventBus
var deleteSuperlikePostService *unsuperlike_post.DeleteSuperlikePostService

func setUpService(t *testing.T) {
	setUp(t)
	serviceRepository = mock_unsuperlike_post.NewMockRepository(ctrl)
	serviceExternalBus = mock_bus.NewMockExternalBus(ctrl)
	serviceBus = bus.NewEventBus(serviceExternalBus)
	deleteSuperlikePostService = unsuperlike_post.NewDeleteSuperlikePostService(serviceRepository, serviceBus)
}

func TestDeleteSuperlikePostWithService_WhenItReturnsSuccess(t *testing.T) {
	setUpService(t)
	superlike := &model.SuperlikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	expectedUserUnsuperlikedPostEvent := &event.UserUnsuperlikedPostEvent{
		Username: superlike.Username,
		PostId:   superlike.PostId,
	}
	expectedEvent, _ := createEvent(event.UserUnsuperlikedPostEventName, expectedUserUnsuperlikedPostEvent)
	serviceRepository.EXPECT().DeleteSuperlikePost(superlike).Return(nil)
	serviceExternalBus.EXPECT().Publish(expectedEvent).Return(nil)

	err := deleteSuperlikePostService.DeleteSuperlikePost(superlike)

	assert.Nil(t, err)
	assert.Contains(t, loggerOutput.String(), fmt.Sprintf("User unsuperliked post, username: %s -> postId: %s", superlike.Username, superlike.PostId))
}

func TestErrorOnDeleteSuperlikePostWithService_WhenDeletingInRepositoryFails(t *testing.T) {
	setUpService(t)
	superlike := &model.SuperlikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	serviceRepository.EXPECT().DeleteSuperlikePost(superlike).Return(errors.New("some error"))

	err := deleteSuperlikePostService.DeleteSuperlikePost(superlike)

	assert.NotNil(t, err)
	assert.Contains(t, loggerOutput.String(), fmt.Sprintf("Error deleting superlike, username: %s -> postId: %s", superlike.Username, superlike.PostId))
}

func TestErrorOnDeleteSuperlikePostWithService_WhenPublishingEventFails(t *testing.T) {
	setUpService(t)
	superlike := &model.SuperlikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	expectedUserUnsuperlikedPostEvent := &event.UserUnsuperlikedPostEvent{
		Username: superlike.Username,
		PostId:   superlike.PostId,
	}
	expectedEvent, _ := createEvent(event.UserUnsuperlikedPostEventName, expectedUserUnsuperlikedPostEvent)
	serviceRepository.EXPECT().DeleteSuperlikePost(superlike).Return(nil)
	serviceExternalBus.EXPECT().Publish(expectedEvent).Return(errors.New("some error"))

	err := deleteSuperlikePostService.DeleteSuperlikePost(superlike)

	assert.NotNil(t, err)
	assert.Contains(t, loggerOutput.String(), fmt.Sprintf("Publishing %s failed, username: %s -> postId: %s", event.UserUnsuperlikedPostEventName, expectedUserUnsuperlikedPostEvent.Username, expectedUserUnsuperlikedPostEvent.PostId))
}
