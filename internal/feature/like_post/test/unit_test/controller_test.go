package unit_test_like_post

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"reactionservice/internal/bus"
	"reactionservice/internal/feature/like_post"
	mock_like_post "reactionservice/internal/feature/like_post/test/mock"

	"github.com/gin-gonic/gin"
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

func TestLikePost(t *testing.T) {
	setUpController(t)
	ginContext.Request = httptest.NewRequest(http.MethodDelete, "/like", nil)
	expectedPostId := uint64(1234)
	postId := strconv.FormatUint(expectedPostId, 10)
	ginContext.Params = []gin.Param{{Key: "postId", Value: postId}}
	controllerService.EXPECT().LikePost(expectedPostId).Return(nil)
	expectedBodyResponse := `{
		"error": false,
		"message": "200 OK",
		"content": null
	}`

	controller.LikePost(ginContext)

	assert.Equal(t, 200, apiResponse.Code)
	assert.Equal(t, removeSpace(apiResponse.Body.String()), removeSpace(expectedBodyResponse))
}

func TestBadRequesErrorOnLikePost_WhenMissingPostId(t *testing.T) {
	setUpController(t)
	ginContext.Request, _ = http.NewRequest("POST", "/like", nil)
	expectedBodyResponse := `{
		"error": true,
		"message": "Missing postId parameter",
		"content": null
	}`

	controller.LikePost(ginContext)

	assert.Equal(t, apiResponse.Code, 400)
	assert.Equal(t, removeSpace(apiResponse.Body.String()), removeSpace(expectedBodyResponse))
}

func TestBadRequesErrorOnLikePost_WhenPostIdIsNotUint64(t *testing.T) {
	setUpController(t)
	ginContext.Request, _ = http.NewRequest("POST", "/like", nil)
	expectedPostId := "no uint64"
	ginContext.Params = []gin.Param{{Key: "postId", Value: expectedPostId}}
	expectedBodyResponse := `{
		"error": true,
		"message": "PostId couldn't be parsed. PostId should be a positive number",
		"content": null
	}`

	controller.LikePost(ginContext)

	assert.Equal(t, apiResponse.Code, 400)
	assert.Equal(t, removeSpace(apiResponse.Body.String()), removeSpace(expectedBodyResponse))
	assert.Contains(t, loggerOutput.String(), fmt.Sprintf("PostId %v couldn't be parsed", expectedPostId))
}

func TestInternalServerErrorOnLikePost(t *testing.T) {
	setUpController(t)
	ginContext.Request = httptest.NewRequest(http.MethodDelete, "/like", nil)
	expectedPostId := uint64(1234)
	ginContext.Params = []gin.Param{{Key: "postId", Value: strconv.FormatUint(expectedPostId, 10)}}
	expectedError := errors.New("some error")
	controllerService.EXPECT().LikePost(expectedPostId).Return(expectedError)
	expectedBodyResponse := `{
		"error": true,
		"message": "` + expectedError.Error() + `",
		"content": null
	}`

	controller.LikePost(ginContext)

	assert.Equal(t, apiResponse.Code, 500)
	assert.Equal(t, removeSpace(apiResponse.Body.String()), removeSpace(expectedBodyResponse))
}
