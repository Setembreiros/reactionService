package unit_test_like_post

import (
	"errors"
	"fmt"
	"testing"

	"reactionservice/internal/bus"
	mock_bus "reactionservice/internal/bus/test/mock"
	"reactionservice/internal/feature/like_post"
	mock_like_post "reactionservice/internal/feature/like_post/test/mock"
	model "reactionservice/internal/model/domain"
	"reactionservice/internal/model/event"

	"github.com/stretchr/testify/assert"
)

var serviceRepository *mock_like_post.MockRepository
var serviceExternalBus *mock_bus.MockExternalBus
var serviceBus *bus.EventBus
var createLikePostService *like_post.LikePostService

func setUpService(t *testing.T) {
	setUp(t)
	serviceRepository = mock_like_post.NewMockRepository(ctrl)
	serviceExternalBus = mock_bus.NewMockExternalBus(ctrl)
	serviceBus = bus.NewEventBus(serviceExternalBus)
	createLikePostService = like_post.NewLikePostService(serviceRepository, serviceBus)
}

func TestCreateLikePostWithService_WhenItReturnsSuccess(t *testing.T) {
	setUpService(t)
	like := &model.LikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	expectedUserLikedPostEvent := &event.UserLikedPostEvent{
		Username: like.Username,
		PostId:   like.PostId,
	}
	expectedEvent, _ := createEvent(event.UserLikedPostEventName, expectedUserLikedPostEvent)
	serviceRepository.EXPECT().CreateLikePost(like).Return(nil)
	serviceExternalBus.EXPECT().Publish(expectedEvent).Return(nil)

	err := createLikePostService.CreateLikePost(like)

	assert.Nil(t, err)
	assert.Contains(t, loggerOutput.String(), fmt.Sprintf("User liked post, username: %s -> postId: %s", like.Username, like.PostId))
}

func TestErrorOnCreateLikePostWithService_WhenAddingToRepositoryFails(t *testing.T) {
	setUpService(t)
	like := &model.LikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	serviceRepository.EXPECT().CreateLikePost(like).Return(errors.New("some error"))

	err := createLikePostService.CreateLikePost(like)

	assert.NotNil(t, err)
	assert.Contains(t, loggerOutput.String(), fmt.Sprintf("Error creating like, username: %s -> postId: %s", like.Username, like.PostId))
}

func TestErrorOnCreateLikePostWithService_WhenPublishingEventFails(t *testing.T) {
	setUpService(t)
	like := &model.LikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	expectedUserLikedPostEvent := &event.UserLikedPostEvent{
		Username: like.Username,
		PostId:   like.PostId,
	}
	expectedEvent, _ := createEvent(event.UserLikedPostEventName, expectedUserLikedPostEvent)
	serviceRepository.EXPECT().CreateLikePost(like).Return(nil)
	serviceExternalBus.EXPECT().Publish(expectedEvent).Return(errors.New("some error"))

	err := createLikePostService.CreateLikePost(like)

	assert.NotNil(t, err)
	assert.Contains(t, loggerOutput.String(), fmt.Sprintf("Publishing %s failed, username: %s -> postId: %s", event.UserLikedPostEventName, expectedUserLikedPostEvent.Username, expectedUserLikedPostEvent.PostId))
}
