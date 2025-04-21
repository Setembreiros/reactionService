package unlike_post

import (
	database "reactionservice/internal/db"
	model "reactionservice/internal/model/domain"
)

type DeleteLikePostRepository struct {
	dataRepository *database.Database
}

func NewDeleteLikePostRepository(dataRepository *database.Database) *DeleteLikePostRepository {
	return &DeleteLikePostRepository{
		dataRepository: dataRepository,
	}
}

func (r *DeleteLikePostRepository) DeleteLikePost(like *model.LikePost) error {
	return r.dataRepository.Client.DeleteLikePost(like)
}
