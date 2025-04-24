package unlike_post

import (
	"reactionservice/internal/bus"
	model "reactionservice/internal/model/domain"
	"reactionservice/internal/model/event"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=service.go -destination=test/mock/service.go

type Repository interface {
	DeleteLikePost(like *model.LikePost) error
}

type DeleteLikePostService struct {
	repository Repository
	bus        *bus.EventBus
}

func NewDeleteLikePostService(repository Repository, bus *bus.EventBus) *DeleteLikePostService {
	return &DeleteLikePostService{
		repository: repository,
		bus:        bus,
	}
}

func (s *DeleteLikePostService) DeleteLikePost(like *model.LikePost) error {
	err := s.repository.DeleteLikePost(like)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Error deleting like, username: %s -> postId: %s", like.Username, like.PostId)
		return err
	}

	err = s.publishUserUnlikedPostEvent(like)
	if err != nil {
		return err
	}

	log.Info().Msgf("User unliked post, username: %s -> postId: %s", like.Username, like.PostId)

	return nil
}

func (s *DeleteLikePostService) publishUserUnlikedPostEvent(like *model.LikePost) error {
	userUnlikedPostEvent := &event.UserUnlikedPostEvent{
		PostId:   like.PostId,
		Username: like.Username,
	}

	err := s.bus.Publish(event.UserUnlikedPostEventName, userUnlikedPostEvent)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Publishing %s failed, username: %s -> postId: %s", event.UserUnlikedPostEventName, userUnlikedPostEvent.Username, userUnlikedPostEvent.PostId)
		return err
	}

	return nil
}
