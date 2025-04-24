package unit_test_superlike_post

import (
	"errors"
	"fmt"
	"testing"

	"reactionservice/internal/bus"
	mock_bus "reactionservice/internal/bus/test/mock"
	"reactionservice/internal/feature/superlike_post"
	mock_superlike_post "reactionservice/internal/feature/superlike_post/test/mock"
	model "reactionservice/internal/model/domain"
	"reactionservice/internal/model/event"

	"github.com/stretchr/testify/assert"
)

var serviceRepository *mock_superlike_post.MockRepository
var serviceExternalBus *mock_bus.MockExternalBus
var serviceBus *bus.EventBus
var createSuperlikePostService *superlike_post.SuperlikePostService

func setUpService(t *testing.T) {
	setUp(t)
	serviceRepository = mock_superlike_post.NewMockRepository(ctrl)
	serviceExternalBus = mock_bus.NewMockExternalBus(ctrl)
	serviceBus = bus.NewEventBus(serviceExternalBus)
	createSuperlikePostService = superlike_post.NewSuperlikePostService(serviceRepository, serviceBus)
}

func TestCreateSuperlikePostWithService_WhenItReturnsSuccess(t *testing.T) {
	setUpService(t)
	superlike := &model.SuperlikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	expectedUserSuperlikedPostEvent := &event.UserSuperlikedPostEvent{
		Username: superlike.Username,
		PostId:   superlike.PostId,
	}
	expectedEvent, _ := createEvent(event.UserSuperlikedPostEventName, expectedUserSuperlikedPostEvent)
	serviceRepository.EXPECT().CreateSuperlikePost(superlike).Return(nil)
	serviceExternalBus.EXPECT().Publish(expectedEvent).Return(nil)

	err := createSuperlikePostService.CreateSuperlikePost(superlike)

	assert.Nil(t, err)
	assert.Contains(t, loggerOutput.String(), fmt.Sprintf("User superliked post, username: %s -> postId: %s", superlike.Username, superlike.PostId))
}

func TestErrorOnCreateSuperlikePostWithService_WhenAddingToRepositoryFails(t *testing.T) {
	setUpService(t)
	superlike := &model.SuperlikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	serviceRepository.EXPECT().CreateSuperlikePost(superlike).Return(errors.New("some error"))

	err := createSuperlikePostService.CreateSuperlikePost(superlike)

	assert.NotNil(t, err)
	assert.Contains(t, loggerOutput.String(), fmt.Sprintf("Error creating superlike, username: %s -> postId: %s", superlike.Username, superlike.PostId))
}

func TestErrorOnCreateSuperlikePostWithService_WhenPublishingEventFails(t *testing.T) {
	setUpService(t)
	superlike := &model.SuperlikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	expectedUserSuperlikedPostEvent := &event.UserSuperlikedPostEvent{
		Username: superlike.Username,
		PostId:   superlike.PostId,
	}
	expectedEvent, _ := createEvent(event.UserSuperlikedPostEventName, expectedUserSuperlikedPostEvent)
	serviceRepository.EXPECT().CreateSuperlikePost(superlike).Return(nil)
	serviceExternalBus.EXPECT().Publish(expectedEvent).Return(errors.New("some error"))

	err := createSuperlikePostService.CreateSuperlikePost(superlike)

	assert.NotNil(t, err)
	assert.Contains(t, loggerOutput.String(), fmt.Sprintf("Publishing %s failed, username: %s -> postId: %s", event.UserSuperlikedPostEventName, expectedUserSuperlikedPostEvent.Username, expectedUserSuperlikedPostEvent.PostId))
}
