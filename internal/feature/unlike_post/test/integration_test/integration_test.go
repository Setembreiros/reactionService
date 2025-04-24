package integration_test_get_unlike_post

import (
	"net/http"
	"net/http/httptest"
	"reactionservice/internal/bus"
	mock_bus "reactionservice/internal/bus/test/mock"
	database "reactionservice/internal/db"
	"reactionservice/internal/feature/unlike_post"
	model "reactionservice/internal/model/domain"
	"reactionservice/internal/model/event"
	integration_test_arrange "reactionservice/test/integration_test_common/arrange"
	integration_test_assert "reactionservice/test/integration_test_common/assert"
	integration_test_builder "reactionservice/test/integration_test_common/builder"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

var serviceExternalBus *mock_bus.MockExternalBus
var db *database.Database
var controller *unlike_post.DeleteLikePostController
var apiResponse *httptest.ResponseRecorder
var ginContext *gin.Context

func setUp(t *testing.T) {
	// Mocks
	gin.SetMode(gin.TestMode)
	apiResponse = httptest.NewRecorder()
	ginContext, _ = gin.CreateTestContext(apiResponse)
	ctrl := gomock.NewController(t)
	serviceExternalBus = mock_bus.NewMockExternalBus(ctrl)
	serviceBus := bus.NewEventBus(serviceExternalBus)

	// Real infrastructure and services
	db = integration_test_arrange.CreateTestDatabase(ginContext)
	repository := unlike_post.NewDeleteLikePostRepository(db)
	service := unlike_post.NewDeleteLikePostService(repository, serviceBus)
	controller = unlike_post.NewDeleteLikePostController(service)
}

func tearDown() {
	db.Client.Clean()
}

func TestDeleteLikePost_WhenDatabaseReturnsSuccess(t *testing.T) {
	setUp(t)
	defer tearDown()
	like := &model.LikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	ginContext.Request = httptest.NewRequest(http.MethodDelete, "/likePost", nil)
	ginContext.Params = []gin.Param{
		{Key: "postId", Value: like.PostId},
		{Key: "username", Value: like.Username},
	}
	expectedBodyResponse := `{
		"error": false,
		"message": "200 OK",
		"content": null
	}`
	expectedUserUnlikedPostEvent := &event.UserUnlikedPostEvent{
		Username: like.Username,
		PostId:   like.PostId,
	}
	expectedEvent := integration_test_builder.NewEventBuilder(t).WithName(event.UserUnlikedPostEventName).WithData(expectedUserUnlikedPostEvent).Build()
	serviceExternalBus.EXPECT().Publish(expectedEvent).Return(nil)

	controller.DeleteLikePost(ginContext)

	integration_test_assert.AssertSuccessResult(t, apiResponse, expectedBodyResponse)
	integration_test_assert.AssertLIkePostDoesNotExists(t, db, like)
}
