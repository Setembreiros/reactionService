package create_review

import (
	"reactionservice/internal/bus"
	model "reactionservice/internal/model/domain"
	"reactionservice/internal/model/event"
	"time"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=service.go -destination=test/mock/service.go

type Repository interface {
	CreateReview(data *model.Review) (uint64, error)
}

type TimeService interface {
	GetTimeNowUtc() time.Time
}

type CreateReviewService struct {
	timeService TimeService
	repository  Repository
	bus         *bus.EventBus
}

func NewCreateReviewService(timeService TimeService, repository Repository, bus *bus.EventBus) *CreateReviewService {
	return &CreateReviewService{
		timeService: timeService,
		repository:  repository,
		bus:         bus,
	}
}

func (s *CreateReviewService) CreateReview(review *model.Review) error {
	review.CreatedAt = s.timeService.GetTimeNowUtc()
	var err error
	review.Id, err = s.repository.CreateReview(review)

	if err != nil {
		log.Error().Stack().Err(err).Msgf("Error creating review, username: %s -> postId: %s", review.Username, review.PostId)
		return err
	}

	err = s.publishReviewWasCreatedEvent(review)
	if err != nil {
		return err
	}

	log.Info().Msgf("Review was created, username: %s -> postId: %s", review.Username, review.PostId)

	return nil
}

func (s *CreateReviewService) publishReviewWasCreatedEvent(data *model.Review) error {
	reviewWasCreatedEvent := &event.ReviewWasCreatedEvent{
		ReviewId:  data.Id,
		Username:  data.Username,
		PostId:    data.PostId,
		Title:     data.Title,
		Content:   data.Content,
		Rating:    data.Rating,
		CreatedAt: data.CreatedAt.Format(model.TimeLayout),
	}

	err := s.bus.Publish(event.ReviewWasCreatedEventName, reviewWasCreatedEvent)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Publishing %s failed, username: %s -> postId: %s", event.ReviewWasCreatedEventName, reviewWasCreatedEvent.Username, reviewWasCreatedEvent.PostId)
		return err
	}

	return nil
}
