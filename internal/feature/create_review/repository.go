package create_review

import (
	database "reactionservice/internal/db"
	model "reactionservice/internal/model/domain"
)

type CreateReviewRepository struct {
	dataRepository *database.Database
}

func NewCreateReviewRepository(dataRepository *database.Database) *CreateReviewRepository {
	return &CreateReviewRepository{
		dataRepository: dataRepository,
	}
}

func (r *CreateReviewRepository) CreateReview(data *model.Review) (uint64, error) {
	return r.dataRepository.Client.CreateReview(data)
}
