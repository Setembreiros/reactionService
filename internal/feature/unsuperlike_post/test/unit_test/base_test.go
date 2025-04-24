package unit_test_unsuperlike_post

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"reactionservice/internal/bus"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog/log"
)

var ctrl *gomock.Controller
var loggerOutput bytes.Buffer
var apiResponse *httptest.ResponseRecorder
var ginContext *gin.Context
var timeLayout string = "2006-01-02T15:04:05.00Z"

func setUp(t *testing.T) {
	ctrl = gomock.NewController(t)
	log.Logger = log.Output(&loggerOutput)
	gin.SetMode(gin.TestMode)
	apiResponse = httptest.NewRecorder()
	ginContext, _ = gin.CreateTestContext(apiResponse)
}

func createEvent(eventName string, eventData any) (*bus.Event, error) {
	dataEvent, err := serializeData(eventData)
	if err != nil {
		return nil, err
	}

	return &bus.Event{
		Type: eventName,
		Data: dataEvent,
	}, nil
}

func serializeData(data any) ([]byte, error) {
	return json.Marshal(data)
}

func removeSpace(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, " ", ""), "\t", ""), "\n", "")
}
