package unit_test_unsuperlike_post

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"reactionservice/internal/bus"
	"reactionservice/internal/feature/unsuperlike_post"
	mock_unsuperlike_post "reactionservice/internal/feature/unsuperlike_post/test/mock"
	model "reactionservice/internal/model/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var controllerLoggerOutput bytes.Buffer
var controllerService *mock_unsuperlike_post.MockService
var controllerBus *bus.EventBus
var controller *unsuperlike_post.DeleteSuperlikePostController

func setUpController(t *testing.T) {
	setUp(t)
	controllerService = mock_unsuperlike_post.NewMockService(ctrl)
	controllerBus = &bus.EventBus{}
	controller = unsuperlike_post.NewDeleteSuperlikePostController(controllerService)
}

func TestCreateSuperlikePost(t *testing.T) {
	setUpController(t)
	superlike := &model.SuperlikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	ginContext.Request = httptest.NewRequest(http.MethodDelete, "/superlikePost", nil)
	ginContext.Params = []gin.Param{
		{Key: "postId", Value: superlike.PostId},
		{Key: "username", Value: superlike.Username},
	}
	controllerService.EXPECT().DeleteSuperlikePost(superlike).Return(nil)
	expectedBodyResponse := `{
		"error": false,
		"message": "200 OK",
		"content": null
	}`

	controller.DeleteSuperlikePost(ginContext)

	assert.Equal(t, 200, apiResponse.Code)
	assert.Equal(t, removeSpace(apiResponse.Body.String()), removeSpace(expectedBodyResponse))
}

func TestBadRequesErrorOnDeleteComment_WhenMissingPostId(t *testing.T) {
	setUpController(t)
	ginContext.Request, _ = http.NewRequest(http.MethodDelete, "/superlikePost", nil)
	ginContext.Params = []gin.Param{
		{Key: "username", Value: "aaaa"},
	}
	expectedBodyResponse := `{
		"error": true,
		"message": "Missing postId parameter",
		"content": null
	}`

	controller.DeleteSuperlikePost(ginContext)

	assert.Equal(t, apiResponse.Code, 400)
	assert.Equal(t, removeSpace(apiResponse.Body.String()), removeSpace(expectedBodyResponse))
}

func TestBadRequesErrorOnDeleteComment_WhenMissingUsername(t *testing.T) {
	setUpController(t)
	ginContext.Request, _ = http.NewRequest(http.MethodDelete, "/superlikePost", nil)
	ginContext.Params = []gin.Param{
		{Key: "postId", Value: "post1"},
	}
	expectedBodyResponse := `{
		"error": true,
		"message": "Missing username parameter",
		"content": null
	}`

	controller.DeleteSuperlikePost(ginContext)

	assert.Equal(t, apiResponse.Code, 400)
	assert.Equal(t, removeSpace(apiResponse.Body.String()), removeSpace(expectedBodyResponse))
}

func TestInternalServerErrorOnDeleteSuperlikePost(t *testing.T) {
	setUpController(t)
	superlike := &model.SuperlikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	ginContext.Request = httptest.NewRequest(http.MethodPost, "/superlikePost", nil)
	ginContext.Params = []gin.Param{
		{Key: "postId", Value: superlike.PostId},
		{Key: "username", Value: superlike.Username},
	}
	expectedError := errors.New("some error")
	controllerService.EXPECT().DeleteSuperlikePost(superlike).Return(expectedError)
	expectedBodyResponse := `{
		"error": true,
		"message": "` + expectedError.Error() + `",
		"content": null
	}`

	controller.DeleteSuperlikePost(ginContext)

	assert.Equal(t, apiResponse.Code, 500)
	assert.Equal(t, removeSpace(apiResponse.Body.String()), removeSpace(expectedBodyResponse))
}
