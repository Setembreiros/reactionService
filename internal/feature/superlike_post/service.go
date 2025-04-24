package superlike_post

import (
	"reactionservice/internal/bus"
	model "reactionservice/internal/model/domain"
	"reactionservice/internal/model/event"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=service.go -destination=test/mock/service.go

type Repository interface {
	CreateSuperlikePost(superlike *model.SuperlikePost) error
}

type SuperlikePostService struct {
	repository Repository
	bus        *bus.EventBus
}

func NewSuperlikePostService(repository Repository, bus *bus.EventBus) *SuperlikePostService {
	return &SuperlikePostService{
		repository: repository,
		bus:        bus,
	}
}

func (s *SuperlikePostService) CreateSuperlikePost(superlike *model.SuperlikePost) error {
	err := s.repository.CreateSuperlikePost(superlike)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Error creating superlike, username: %s -> postId: %s", superlike.Username, superlike.PostId)
		return err
	}

	err = s.publishUserSuperlikedPostEvent(superlike)
	if err != nil {
		return err
	}

	log.Info().Msgf("User superliked post, username: %s -> postId: %s", superlike.Username, superlike.PostId)

	return nil
}

func (s *SuperlikePostService) publishUserSuperlikedPostEvent(superlike *model.SuperlikePost) error {
	userSuperlikedPostEvent := &event.UserSuperlikedPostEvent{
		PostId:   superlike.PostId,
		Username: superlike.Username,
	}

	err := s.bus.Publish(event.UserSuperlikedPostEventName, userSuperlikedPostEvent)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Publishing %s failed, username: %s -> postId: %s", event.UserSuperlikedPostEventName, userSuperlikedPostEvent.Username, userSuperlikedPostEvent.PostId)
		return err
	}

	return nil
}
