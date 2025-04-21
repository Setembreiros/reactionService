package like_post

import (
	"reactionservice/internal/bus"
	model "reactionservice/internal/model/domain"
	"reactionservice/internal/model/event"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=service.go -destination=test/mock/service.go

type Repository interface {
	CreateLikePost(like *model.LikePost) error
}

type LikePostService struct {
	repository Repository
	bus        *bus.EventBus
}

func NewLikePostService(repository Repository, bus *bus.EventBus) *LikePostService {
	return &LikePostService{
		repository: repository,
		bus:        bus,
	}
}

func (s *LikePostService) CreateLikePost(like *model.LikePost) error {
	err := s.repository.CreateLikePost(like)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Error creating like, username: %s -> postId: %s", like.Username, like.PostId)
		return err
	}

	err = s.publishUserLikedPostEvent(like)
	if err != nil {
		return err
	}

	log.Info().Msgf("User liked post, username: %s -> postId: %s", like.Username, like.PostId)

	return nil
}

func (s *LikePostService) publishUserLikedPostEvent(like *model.LikePost) error {
	userLikedPostEvent := &event.UserLikedPostEvent{
		PostId:   like.PostId,
		Username: like.Username,
	}

	err := s.bus.Publish(event.UserLikedPostEventName, userLikedPostEvent)
	if err != nil {
		log.Error().Stack().Err(err).Msgf("Publishing %s failed, username: %s -> postId: %s", event.UserLikedPostEventName, userLikedPostEvent.Username, userLikedPostEvent.PostId)
		return err
	}

	return nil
}
