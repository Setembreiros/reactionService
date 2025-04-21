package unit_test_like_post

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"reactionservice/internal/bus"
	"reactionservice/internal/feature/like_post"
	mock_like_post "reactionservice/internal/feature/like_post/test/mock"
	model "reactionservice/internal/model/domain"

	"github.com/stretchr/testify/assert"
)

var controllerLoggerOutput bytes.Buffer
var controllerService *mock_like_post.MockService
var controllerBus *bus.EventBus
var controller *like_post.LikePostController

func setUpController(t *testing.T) {
	setUp(t)
	controllerService = mock_like_post.NewMockService(ctrl)
	controllerBus = &bus.EventBus{}
	controller = like_post.NewLikePostController(controllerService)
}

func TestCreateLikePost(t *testing.T) {
	setUpController(t)
	like := &model.LikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	data, _ := serializeData(like)
	ginContext.Request = httptest.NewRequest(http.MethodPost, "/likePost", bytes.NewBuffer(data))
	controllerService.EXPECT().CreateLikePost(like).Return(nil)
	expectedBodyResponse := `{
		"error": false,
		"message": "200 OK",
		"content": null
	}`

	controller.CreateLikePost(ginContext)

	assert.Equal(t, 200, apiResponse.Code)
	assert.Equal(t, removeSpace(apiResponse.Body.String()), removeSpace(expectedBodyResponse))
}

func TestBadRequesErrorOnCreateLikePost_WhenInvalidData(t *testing.T) {
	setUpController(t)
	invalidData := "invalid data"
	ginContext.Request, _ = http.NewRequest(http.MethodPost, "/likePost", bytes.NewBuffer([]byte(invalidData)))
	expectedBodyResponse := `{
		"error": true,
		"message": "Invalid Json Request",
		"content": null
	}`

	controller.CreateLikePost(ginContext)

	assert.Equal(t, apiResponse.Code, 400)
	assert.Equal(t, removeSpace(apiResponse.Body.String()), removeSpace(expectedBodyResponse))
}

func TestInternalServerErrorOnLikePost(t *testing.T) {
	setUpController(t)
	like := &model.LikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	data, _ := serializeData(like)
	ginContext.Request = httptest.NewRequest(http.MethodPost, "/likePost", bytes.NewBuffer(data))
	expectedError := errors.New("some error")
	controllerService.EXPECT().CreateLikePost(like).Return(expectedError)
	expectedBodyResponse := `{
		"error": true,
		"message": "` + expectedError.Error() + `",
		"content": null
	}`

	controller.CreateLikePost(ginContext)

	assert.Equal(t, apiResponse.Code, 500)
	assert.Equal(t, removeSpace(apiResponse.Body.String()), removeSpace(expectedBodyResponse))
}
