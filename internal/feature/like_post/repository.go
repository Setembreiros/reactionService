package like_post

import (
	database "reactionservice/internal/db"
	model "reactionservice/internal/model/domain"
	customerror "reactionservice/internal/model/error"
)

type CreateLikePostRepository struct {
	dataRepository *database.Database
}

func NewCreateLikePostRepository(dataRepository *database.Database) *CreateLikePostRepository {
	return &CreateLikePostRepository{
		dataRepository: dataRepository,
	}
}

func (r *CreateLikePostRepository) CreateLikePost(like *model.LikePost) error {
	likeExists, err := r.dataRepository.Client.GetLikePost(like.PostId, like.Username)
	if err != nil {
		return err
	}
	if likeExists != nil {
		return customerror.NewDataAlreadyExistsError("Like")
	}
	return r.dataRepository.Client.CreateLikePost(like)
}
