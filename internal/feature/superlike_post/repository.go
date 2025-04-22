package superlike_post

import (
	database "reactionservice/internal/db"
	model "reactionservice/internal/model/domain"
	customerror "reactionservice/internal/model/error"
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
	superlikeExists, err := r.dataRepository.Client.GetSuperlikePost(superlike.PostId, superlike.Username)
	if err != nil {
		return err
	}
	if superlikeExists != nil {
		return customerror.NewDataAlreadyExistsError("Superlike")
	}
	return r.dataRepository.Client.CreateSuperlikePost(superlike)
}
