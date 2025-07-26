package unit_test_create_review

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"reactionservice/internal/bus"
	mock_bus "reactionservice/internal/bus/test/mock"
	"reactionservice/internal/feature/create_review"
	mock_create_review "reactionservice/internal/feature/create_review/test/mock"
	model "reactionservice/internal/model/domain"
	"reactionservice/internal/model/event"

	"github.com/stretchr/testify/assert"
)

var timeService *mock_create_review.MockTimeService
var serviceRepository *mock_create_review.MockRepository
var serviceExternalBus *mock_bus.MockExternalBus
var serviceBus *bus.EventBus
var createReviewService *create_review.CreateReviewService

func setUpService(t *testing.T) {
	setUp(t)
	timeService = mock_create_review.NewMockTimeService(ctrl)
	serviceRepository = mock_create_review.NewMockRepository(ctrl)
	serviceExternalBus = mock_bus.NewMockExternalBus(ctrl)
	serviceBus = bus.NewEventBus(serviceExternalBus)
	createReviewService = create_review.NewCreateReviewService(timeService, serviceRepository, serviceBus)
}

func TestCreateReviewWithService_WhenItReturnsSuccess(t *testing.T) {
	setUpService(t)
	review := &model.Review{
		Username: "usernameA",
		PostId:   "post1",
		Title:    "Título da miña review",
		Content:  "a minha review",
		Rating:   2,
	}
	timeNowString := time.Now().UTC().Format(model.TimeLayout)
	timeNow, _ := time.Parse(model.TimeLayout, timeNowString)
	review.CreatedAt = timeNow
	expectedReviewId := uint64(1000)
	expectedReviewWasCreatedEvent := &event.ReviewWasCreatedEvent{
		ReviewId:  expectedReviewId,
		Username:  review.Username,
		PostId:    review.PostId,
		Title:     review.Title,
		Content:   review.Content,
		Rating:    review.Rating,
		CreatedAt: review.CreatedAt.Format(model.TimeLayout),
	}
	expectedEvent, _ := createEvent(event.ReviewWasCreatedEventName, expectedReviewWasCreatedEvent)
	timeService.EXPECT().GetTimeNowUtc().Return(timeNow)
	serviceRepository.EXPECT().CreateReview(review).Return(expectedReviewId, nil)
	serviceExternalBus.EXPECT().Publish(expectedEvent).Return(nil)

	err := createReviewService.CreateReview(review)

	assert.Nil(t, err)
	assert.Contains(t, loggerOutput.String(), fmt.Sprintf("Review was created, username: %s -> postId: %s", review.Username, review.PostId))
}

func TestErrorOnCreateReviewWithService_WhenAddingToRepositoryFails(t *testing.T) {
	setUpService(t)
	review := &model.Review{
		Username: "usernameA",
		PostId:   "post1",
		Title:    "Título da miña review",
		Content:  "a miña review",
		Rating:   2,
	}
	timeNowString := time.Now().UTC().Format(model.TimeLayout)
	timeNow, _ := time.Parse(model.TimeLayout, timeNowString)
	review.CreatedAt = timeNow
	timeService.EXPECT().GetTimeNowUtc().Return(timeNow)
	serviceRepository.EXPECT().CreateReview(review).Return(uint64(0), errors.New("some error"))

	err := createReviewService.CreateReview(review)

	assert.NotNil(t, err)
	assert.Contains(t, loggerOutput.String(), fmt.Sprintf("Error creating review, username: %s -> postId: %s", review.Username, review.PostId))
}

func TestErrorOnCreateReviewWithService_WhenPublishingEventFails(t *testing.T) {
	setUpService(t)
	review := &model.Review{
		Username: "usernameA",
		PostId:   "post1",
		Title:    "Título da miña review",
		Content:  "a minha review",
		Rating:   2,
	}
	timeNowString := time.Now().UTC().Format(model.TimeLayout)
	timeNow, _ := time.Parse(model.TimeLayout, timeNowString)
	review.CreatedAt = timeNow
	expectedReviewId := uint64(5)
	expectedReviewWasCreatedEvent := &event.ReviewWasCreatedEvent{
		ReviewId:  expectedReviewId,
		Username:  review.Username,
		PostId:    review.PostId,
		Title:     review.Title,
		Content:   review.Content,
		Rating:    review.Rating,
		CreatedAt: review.CreatedAt.Format(model.TimeLayout),
	}
	expectedEvent, _ := createEvent(event.ReviewWasCreatedEventName, expectedReviewWasCreatedEvent)
	timeService.EXPECT().GetTimeNowUtc().Return(timeNow)
	serviceRepository.EXPECT().CreateReview(review).Return(expectedReviewId, nil)
	serviceExternalBus.EXPECT().Publish(expectedEvent).Return(errors.New("some error"))

	err := createReviewService.CreateReview(review)

	assert.NotNil(t, err)
	assert.Contains(t, loggerOutput.String(), fmt.Sprintf("Publishing %s failed, username: %s -> postId: %s", event.ReviewWasCreatedEventName, expectedReviewWasCreatedEvent.Username, expectedReviewWasCreatedEvent.PostId))
}
