package unit_test_unlike_post

import (
	"errors"
	"fmt"
	"testing"

	"reactionservice/internal/bus"
	mock_bus "reactionservice/internal/bus/test/mock"
	"reactionservice/internal/feature/unlike_post"
	mock_unlike_post "reactionservice/internal/feature/unlike_post/test/mock"
	model "reactionservice/internal/model/domain"
	"reactionservice/internal/model/event"

	"github.com/stretchr/testify/assert"
)

var serviceRepository *mock_unlike_post.MockRepository
var serviceExternalBus *mock_bus.MockExternalBus
var serviceBus *bus.EventBus
var deleteLikePostService *unlike_post.DeleteLikePostService

func setUpService(t *testing.T) {
	setUp(t)
	serviceRepository = mock_unlike_post.NewMockRepository(ctrl)
	serviceExternalBus = mock_bus.NewMockExternalBus(ctrl)
	serviceBus = bus.NewEventBus(serviceExternalBus)
	deleteLikePostService = unlike_post.NewDeleteLikePostService(serviceRepository, serviceBus)
}

func TestDeleteLikePostWithService_WhenItReturnsSuccess(t *testing.T) {
	setUpService(t)
	like := &model.LikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	expectedUserUnlikedPostEvent := &event.UserUnlikedPostEvent{
		Username: like.Username,
		PostId:   like.PostId,
	}
	expectedEvent, _ := createEvent(event.UserUnlikedPostEventName, expectedUserUnlikedPostEvent)
	serviceRepository.EXPECT().DeleteLikePost(like).Return(nil)
	serviceExternalBus.EXPECT().Publish(expectedEvent).Return(nil)

	err := deleteLikePostService.DeleteLikePost(like)

	assert.Nil(t, err)
	assert.Contains(t, loggerOutput.String(), fmt.Sprintf("User unliked post, username: %s -> postId: %s", like.Username, like.PostId))
}

func TestErrorOnDeleteLikePostWithService_WhenDeletingInRepositoryFails(t *testing.T) {
	setUpService(t)
	like := &model.LikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	serviceRepository.EXPECT().DeleteLikePost(like).Return(errors.New("some error"))

	err := deleteLikePostService.DeleteLikePost(like)

	assert.NotNil(t, err)
	assert.Contains(t, loggerOutput.String(), fmt.Sprintf("Error deleting like, username: %s -> postId: %s", like.Username, like.PostId))
}

func TestErrorOnDeleteLikePostWithService_WhenPublishingEventFails(t *testing.T) {
	setUpService(t)
	like := &model.LikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	expectedUserUnlikedPostEvent := &event.UserUnlikedPostEvent{
		Username: like.Username,
		PostId:   like.PostId,
	}
	expectedEvent, _ := createEvent(event.UserUnlikedPostEventName, expectedUserUnlikedPostEvent)
	serviceRepository.EXPECT().DeleteLikePost(like).Return(nil)
	serviceExternalBus.EXPECT().Publish(expectedEvent).Return(errors.New("some error"))

	err := deleteLikePostService.DeleteLikePost(like)

	assert.NotNil(t, err)
	assert.Contains(t, loggerOutput.String(), fmt.Sprintf("Publishing %s failed, username: %s -> postId: %s", event.UserUnlikedPostEventName, expectedUserUnlikedPostEvent.Username, expectedUserUnlikedPostEvent.PostId))
}
