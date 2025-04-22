package unsuperlike_post

import (
	database "reactionservice/internal/db"
	model "reactionservice/internal/model/domain"
)

type DeleteSuperlikePostRepository struct {
	dataRepository *database.Database
}

func NewDeleteSuperlikePostRepository(dataRepository *database.Database) *DeleteSuperlikePostRepository {
	return &DeleteSuperlikePostRepository{
		dataRepository: dataRepository,
	}
}

func (r *DeleteSuperlikePostRepository) DeleteSuperlikePost(superlike *model.SuperlikePost) error {
	return r.dataRepository.Client.DeleteSuperlikePost(superlike)
}
