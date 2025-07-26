package integration_test_create_review

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reactionservice/internal/bus"
	mock_bus "reactionservice/internal/bus/test/mock"
	database "reactionservice/internal/db"
	"reactionservice/internal/feature/create_review"
	mock_create_review "reactionservice/internal/feature/create_review/test/mock"
	model "reactionservice/internal/model/domain"
	"reactionservice/internal/model/event"
	integration_test_arrange "reactionservice/test/integration_test_common/arrange"
	integration_test_assert "reactionservice/test/integration_test_common/assert"
	integration_test_builder "reactionservice/test/integration_test_common/builder"
	"reactionservice/test/test_common"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

var timeService *mock_create_review.MockTimeService
var serviceExternalBus *mock_bus.MockExternalBus
var db *database.Database
var controller *create_review.CreateReviewController
var apiResponse *httptest.ResponseRecorder
var ginContext *gin.Context

func setUp(t *testing.T) {
	// Mocks
	gin.SetMode(gin.TestMode)
	apiResponse = httptest.NewRecorder()
	ginContext, _ = gin.CreateTestContext(apiResponse)
	ctrl := gomock.NewController(t)
	timeService = mock_create_review.NewMockTimeService(ctrl)
	serviceExternalBus = mock_bus.NewMockExternalBus(ctrl)
	serviceBus := bus.NewEventBus(serviceExternalBus)

	// Real infrastructure and services
	db = integration_test_arrange.CreateTestDatabase(ginContext)
	repository := create_review.NewCreateReviewRepository(db)
	service := create_review.NewCreateReviewService(timeService, repository, serviceBus)
	controller = create_review.NewCreateReviewController(service)
}

func tearDown() {
	db.Client.Clean()
}

func TestCreateReview_WhenDatabaseReturnsSuccess(t *testing.T) {
	setUp(t)
	defer tearDown()
	review := &model.Review{
		Username: "usernameA",
		PostId:   "post1",
		Content:  "a miña review",
		Rating:   2,
	}
	data, _ := test_common.SerializeData(review)
	ginContext.Request = httptest.NewRequest(http.MethodPost, "/review", bytes.NewBuffer(data))
	expectedBodyResponse := `{
		"error": false,
		"message": "200 OK",
		"content": null
	}`
	timeNowString := time.Now().UTC().Format(model.TimeLayout)
	timeNow, _ := time.Parse(model.TimeLayout, timeNowString)
	timeService.EXPECT().GetTimeNowUtc().Return(timeNow)
	reviewId := integration_test_arrange.GetNextReviewId()
	expectedReview := &model.Review{
		Id:        reviewId,
		Username:  review.Username,
		PostId:    review.PostId,
		Content:   review.Content,
		Rating:    review.Rating,
		CreatedAt: timeNow,
	}
	expectedReviewWasCreatedEvent := &event.ReviewWasCreatedEvent{
		ReviewId:  reviewId,
		Username:  review.Username,
		PostId:    review.PostId,
		Content:   review.Content,
		Rating:    review.Rating,
		CreatedAt: timeNowString,
	}
	expectedEvent := integration_test_builder.NewEventBuilder(t).WithName(event.ReviewWasCreatedEventName).WithData(expectedReviewWasCreatedEvent).Build()
	serviceExternalBus.EXPECT().Publish(expectedEvent).Return(nil)

	controller.CreateReview(ginContext)

	integration_test_assert.AssertSuccessResult(t, apiResponse, expectedBodyResponse)
	integration_test_assert.AssertReviewExists(t, db, reviewId, expectedReview)
}
