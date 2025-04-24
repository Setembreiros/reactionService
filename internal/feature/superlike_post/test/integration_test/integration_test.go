package integration_test_get_superlike_post

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reactionservice/internal/bus"
	mock_bus "reactionservice/internal/bus/test/mock"
	database "reactionservice/internal/db"
	"reactionservice/internal/feature/superlike_post"
	model "reactionservice/internal/model/domain"
	"reactionservice/internal/model/event"
	integration_test_arrange "reactionservice/test/integration_test_common/arrange"
	integration_test_assert "reactionservice/test/integration_test_common/assert"
	integration_test_builder "reactionservice/test/integration_test_common/builder"
	"reactionservice/test/test_common"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

var serviceExternalBus *mock_bus.MockExternalBus
var db *database.Database
var controller *superlike_post.SuperlikePostController
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
	repository := superlike_post.NewCreateSuperlikePostRepository(db)
	service := superlike_post.NewSuperlikePostService(repository, serviceBus)
	controller = superlike_post.NewSuperlikePostController(service)
}

func tearDown() {
	db.Client.Clean()
}

func TestCreateSuperlikePost_WhenDatabaseReturnsSuccess(t *testing.T) {
	setUp(t)
	defer tearDown()
	superlike := &model.SuperlikePost{
		Username: "usernameA",
		PostId:   "post1",
	}
	data, _ := test_common.SerializeData(superlike)
	ginContext.Request = httptest.NewRequest(http.MethodPost, "/superlikePost", bytes.NewBuffer(data))
	expectedBodyResponse := `{
		"error": false,
		"message": "200 OK",
		"content": null
	}`
	expectedUserSuperlikedPostEvent := &event.UserSuperlikedPostEvent{
		Username: superlike.Username,
		PostId:   superlike.PostId,
	}
	expectedEvent := integration_test_builder.NewEventBuilder(t).WithName(event.UserSuperlikedPostEventName).WithData(expectedUserSuperlikedPostEvent).Build()
	serviceExternalBus.EXPECT().Publish(expectedEvent).Return(nil)

	controller.CreateSuperlikePost(ginContext)

	integration_test_assert.AssertSuccessResult(t, apiResponse, expectedBodyResponse)
	integration_test_assert.AssertSuperlikePostExists(t, db, superlike)
}
