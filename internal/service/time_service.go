package service

import (
	model "reactionservice/internal/model/domain"
	"time"
)

type TimeService struct {
}

var timeService TimeService

func GetTimeServiceInstance() *TimeService {
	return &timeService
}

func (t *TimeService) GetTimeNowUtc() time.Time {
	timeNowString := time.Now().UTC().Format(model.TimeLayout)
	timeNow, _ := time.Parse(model.TimeLayout, timeNowString)

	return timeNow
}
