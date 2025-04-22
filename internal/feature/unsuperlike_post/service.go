package unsuperlike_post

import (
	"reactionservice/internal/bus"
	model "reactionservice/internal/model/domain"
	"reactionservice/internal/model/event"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=service.go -destination=test/mock/service.go

type Repository interface {
	DeleteSuperlikePost(superlike *model.SuperlikePost) error
}

type DeleteSuperlikePostService struct {
	repository Repository
	bus        *bus.EventBus
}

func NewDeleteSuperlikePostService(repository Repository, bus *bus.EventBus) *DeleteSuperlikePostService {
	return &DeleteSuperlikePostService{
		repository: repository,
		bus:        bus,
	}
}

func (s *DeleteSuperlikePostService) DeleteSuperlikePost(superlike *model.SuperlikePost) error {
	err := s.repository.DeleteSuperlikePost(superlike)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Error deleting superlike, username: %s -> postId: %s", superlike.Username, superlike.PostId)
		return err
	}

	err = s.publishUserUnsuperlikedPostEvent(superlike)
	if err != nil {
		return err
	}

	log.Info().Msgf("User unsuperliked post, username: %s -> postId: %s", superlike.Username, superlike.PostId)

	return nil
}

func (s *DeleteSuperlikePostService) publishUserUnsuperlikedPostEvent(superlike *model.SuperlikePost) error {
	userUnsuperlikedPostEvent := &event.UserUnsuperlikedPostEvent{
		PostId:   superlike.PostId,
		Username: superlike.Username,
	}

	err := s.bus.Publish(event.UserUnsuperlikedPostEventName, userUnsuperlikedPostEvent)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Publishing %s failed, username: %s -> postId: %s", event.UserUnsuperlikedPostEventName, userUnsuperlikedPostEvent.Username, userUnsuperlikedPostEvent.PostId)
		return err
	}

	return nil
}
