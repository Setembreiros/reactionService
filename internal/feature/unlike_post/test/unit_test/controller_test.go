package unit_test_unlike_post

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"reactionservice/internal/bus"
	"reactionservice/internal/feature/unlike_post"
	mock_unlike_post "reactionservice/internal/feature/unlike_post/test/mock"
	model "reactionservice/internal/model/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var controllerLoggerOutput bytes.Buffer
var controllerService *mock_unlike_post.MockService
var controllerBus *bus.EventBus
var controller *unlike_post.DeleteLikePostController

func setUpController(t *testing.T) {
	setUp(t)
	controllerService = mock_unlike_post.NewMockService(ctrl)
	controllerBus = &bus.EventBus{}
	controller = unlike_post.NewDeleteLikePostController(controllerService)
}

func TestCreateLikePost(t *testing.T) {
	setUpController(t)
	like := &model.LikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	ginContext.Request = httptest.NewRequest(http.MethodDelete, "/likePost", nil)
	ginContext.Params = []gin.Param{
		{Key: "postId", Value: like.PostId},
		{Key: "username", Value: like.Username},
	}
	controllerService.EXPECT().DeleteLikePost(like).Return(nil)
	expectedBodyResponse := `{
		"error": false,
		"message": "200 OK",
		"content": null
	}`

	controller.DeleteLikePost(ginContext)

	assert.Equal(t, 200, apiResponse.Code)
	assert.Equal(t, removeSpace(apiResponse.Body.String()), removeSpace(expectedBodyResponse))
}

func TestBadRequesErrorOnDeleteComment_WhenMissingPostId(t *testing.T) {
	setUpController(t)
	ginContext.Request, _ = http.NewRequest(http.MethodDelete, "/likePost", nil)
	ginContext.Params = []gin.Param{
		{Key: "username", Value: "aaaa"},
	}
	expectedBodyResponse := `{
		"error": true,
		"message": "Missing postId parameter",
		"content": null
	}`

	controller.DeleteLikePost(ginContext)

	assert.Equal(t, apiResponse.Code, 400)
	assert.Equal(t, removeSpace(apiResponse.Body.String()), removeSpace(expectedBodyResponse))
}

func TestBadRequesErrorOnDeleteComment_WhenMissingUsername(t *testing.T) {
	setUpController(t)
	ginContext.Request, _ = http.NewRequest(http.MethodDelete, "/likePost", nil)
	ginContext.Params = []gin.Param{
		{Key: "postId", Value: "post1"},
	}
	expectedBodyResponse := `{
		"error": true,
		"message": "Missing username parameter",
		"content": null
	}`

	controller.DeleteLikePost(ginContext)

	assert.Equal(t, apiResponse.Code, 400)
	assert.Equal(t, removeSpace(apiResponse.Body.String()), removeSpace(expectedBodyResponse))
}

func TestInternalServerErrorOnDeleteLikePost(t *testing.T) {
	setUpController(t)
	like := &model.LikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	ginContext.Request = httptest.NewRequest(http.MethodPost, "/likePost", nil)
	ginContext.Params = []gin.Param{
		{Key: "postId", Value: like.PostId},
		{Key: "username", Value: like.Username},
	}
	expectedError := errors.New("some error")
	controllerService.EXPECT().DeleteLikePost(like).Return(expectedError)
	expectedBodyResponse := `{
		"error": true,
		"message": "` + expectedError.Error() + `",
		"content": null
	}`

	controller.DeleteLikePost(ginContext)

	assert.Equal(t, apiResponse.Code, 500)
	assert.Equal(t, removeSpace(apiResponse.Body.String()), removeSpace(expectedBodyResponse))
}
