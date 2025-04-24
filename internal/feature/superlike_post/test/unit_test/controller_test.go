package unit_test_superlike_post

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"reactionservice/internal/bus"
	"reactionservice/internal/feature/superlike_post"
	mock_superlike_post "reactionservice/internal/feature/superlike_post/test/mock"
	model "reactionservice/internal/model/domain"

	"github.com/stretchr/testify/assert"
)

var controllerLoggerOutput bytes.Buffer
var controllerService *mock_superlike_post.MockService
var controllerBus *bus.EventBus
var controller *superlike_post.SuperlikePostController

func setUpController(t *testing.T) {
	setUp(t)
	controllerService = mock_superlike_post.NewMockService(ctrl)
	controllerBus = &bus.EventBus{}
	controller = superlike_post.NewSuperlikePostController(controllerService)
}

func TestCreateSuperlikePost(t *testing.T) {
	setUpController(t)
	superlike := &model.SuperlikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	data, _ := serializeData(superlike)
	ginContext.Request = httptest.NewRequest(http.MethodPost, "/superlikePost", bytes.NewBuffer(data))
	controllerService.EXPECT().CreateSuperlikePost(superlike).Return(nil)
	expectedBodyResponse := `{
		"error": false,
		"message": "200 OK",
		"content": null
	}`

	controller.CreateSuperlikePost(ginContext)

	assert.Equal(t, 200, apiResponse.Code)
	assert.Equal(t, removeSpace(apiResponse.Body.String()), removeSpace(expectedBodyResponse))
}

func TestBadRequesErrorOnCreateSuperlikePost_WhenInvalidData(t *testing.T) {
	setUpController(t)
	invalidData := "invalid data"
	ginContext.Request, _ = http.NewRequest(http.MethodPost, "/superlikePost", bytes.NewBuffer([]byte(invalidData)))
	expectedBodyResponse := `{
		"error": true,
		"message": "Invalid Json Request",
		"content": null
	}`

	controller.CreateSuperlikePost(ginContext)

	assert.Equal(t, apiResponse.Code, 400)
	assert.Equal(t, removeSpace(apiResponse.Body.String()), removeSpace(expectedBodyResponse))
}

func TestInternalServerErrorOnSuperlikePost(t *testing.T) {
	setUpController(t)
	superlike := &model.SuperlikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	data, _ := serializeData(superlike)
	ginContext.Request = httptest.NewRequest(http.MethodPost, "/superlikePost", bytes.NewBuffer(data))
	expectedError := errors.New("some error")
	controllerService.EXPECT().CreateSuperlikePost(superlike).Return(expectedError)
	expectedBodyResponse := `{
		"error": true,
		"message": "` + expectedError.Error() + `",
		"content": null
	}`

	controller.CreateSuperlikePost(ginContext)

	assert.Equal(t, apiResponse.Code, 500)
	assert.Equal(t, removeSpace(apiResponse.Body.String()), removeSpace(expectedBodyResponse))
}
