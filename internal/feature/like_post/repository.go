package like_post

import (
	database "reactionservice/internal/db"
	model "reactionservice/internal/model/domain"
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
	return r.dataRepository.Client.CreateLikePost(like)
}
