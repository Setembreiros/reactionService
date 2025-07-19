package unit_test_create_review

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"reactionservice/internal/bus"
	"reactionservice/internal/feature/create_review"
	mock_create_review "reactionservice/internal/feature/create_review/test/mock"
	model "reactionservice/internal/model/domain"

	"github.com/go-playground/assert/v2"
)

var controllerLoggerOutput bytes.Buffer
var controllerService *mock_create_review.MockService
var controllerBus *bus.EventBus
var controller *create_review.CreateReviewController

func setUpHandler(t *testing.T) {
	setUp(t)
	controllerService = mock_create_review.NewMockService(ctrl)
	controllerBus = &bus.EventBus{}
	controller = create_review.NewCreateReviewController(controllerService)
}

func TestCreateReview(t *testing.T) {
	setUpHandler(t)
	review := &model.Review{
		Username: "usernameA",
		PostId:   "post1",
		Content:  "a miña review",
		Rating:   2,
	}
	data, _ := serializeData(review)
	ginContext.Request = httptest.NewRequest(http.MethodPost, "/review", bytes.NewBuffer(data))
	controllerService.EXPECT().CreateReview(review).Return(nil)
	expectedBodyResponse := `{
		"error": false,
		"message": "200 OK",
		"content": null
	}`

	controller.CreateReview(ginContext)

	assert.Equal(t, apiResponse.Code, 200)
	assert.Equal(t, removeSpace(apiResponse.Body.String()), removeSpace(expectedBodyResponse))
}

func TestInternalServerErrorOnCreateReview(t *testing.T) {
	setUpHandler(t)
	review := &model.Review{
		Username: "usernameA",
		PostId:   "post1",
		Content:  "a miña review",
		Rating:   2,
	}
	data, _ := serializeData(review)
	ginContext.Request = httptest.NewRequest(http.MethodPost, "/review", bytes.NewBuffer(data))
	expectedError := errors.New("some error")
	controllerService.EXPECT().CreateReview(review).Return(expectedError)
	expectedBodyResponse := `{
		"error": true,
		"message": "` + expectedError.Error() + `",
		"content": null
	}`

	controller.CreateReview(ginContext)

	assert.Equal(t, apiResponse.Code, 500)
	assert.Equal(t, removeSpace(apiResponse.Body.String()), removeSpace(expectedBodyResponse))
}
