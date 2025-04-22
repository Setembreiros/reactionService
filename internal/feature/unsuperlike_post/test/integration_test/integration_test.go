package integration_test_get_unsuperlike_post

import (
	"net/http"
	"net/http/httptest"
	"reactionservice/internal/bus"
	mock_bus "reactionservice/internal/bus/test/mock"
	database "reactionservice/internal/db"
	"reactionservice/internal/feature/unsuperlike_post"
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
var controller *unsuperlike_post.DeleteSuperlikePostController
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
	repository := unsuperlike_post.NewDeleteSuperlikePostRepository(db)
	service := unsuperlike_post.NewDeleteSuperlikePostService(repository, serviceBus)
	controller = unsuperlike_post.NewDeleteSuperlikePostController(service)
}

func tearDown() {
	db.Client.Clean()
}

func TestDeleteSuperlikePost_WhenDatabaseReturnsSuccess(t *testing.T) {
	setUp(t)
	defer tearDown()
	superlike := &model.SuperlikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	integration_test_arrange.AddSuperlikePost(t, superlike)
	ginContext.Request = httptest.NewRequest(http.MethodDelete, "/superlikePost", nil)
	ginContext.Params = []gin.Param{
		{Key: "postId", Value: superlike.PostId},
		{Key: "username", Value: superlike.Username},
	}
	expectedBodyResponse := `{
		"error": false,
		"message": "200 OK",
		"content": null
	}`
	expectedUserUnsuperlikedPostEvent := &event.UserUnsuperlikedPostEvent{
		Username: superlike.Username,
		PostId:   superlike.PostId,
	}
	expectedEvent := integration_test_builder.NewEventBuilder(t).WithName(event.UserUnsuperlikedPostEventName).WithData(expectedUserUnsuperlikedPostEvent).Build()
	serviceExternalBus.EXPECT().Publish(expectedEvent).Return(nil)

	controller.DeleteSuperlikePost(ginContext)

	integration_test_assert.AssertSuccessResult(t, apiResponse, expectedBodyResponse)
	integration_test_assert.AssertSuperlikePostDoesNotExists(t, db, superlike)
}
