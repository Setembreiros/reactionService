package superlike_post

import (
	database "reactionservice/internal/db"
	model "reactionservice/internal/model/domain"
)

type CreateSuperlikePostRepository struct {
	dataRepository *database.Database
}

func NewCreateSuperlikePostRepository(dataRepository *database.Database) *CreateSuperlikePostRepository {
	return &CreateSuperlikePostRepository{
		dataRepository: dataRepository,
	}
}

func (r *CreateSuperlikePostRepository) CreateSuperlikePost(superlike *model.SuperlikePost) error {
	return r.dataRepository.Client.CreateSuperlikePost(superlike)
}
