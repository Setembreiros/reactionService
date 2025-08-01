package database

import (
	model "reactionservice/internal/model/domain"

	_ "github.com/lib/pq"
)

//go:generate mockgen -source=database.go -destination=test/mock/database.go

type Database struct {
	Client DatabaseClient
}

type DatabaseClient interface {
	Clean()
	CreateLikePost(data *model.LikePost) error
	CreateSuperlikePost(data *model.SuperlikePost) error
	CreateReview(data *model.Review) (uint64, error)
	GetLikePost(postId, username string) (*model.LikePost, error)
	GetSuperlikePost(postId, username string) (*model.SuperlikePost, error)
	GetReviewById(id uint64) (*model.Review, error)
	DeleteLikePost(data *model.LikePost) error
	DeleteSuperlikePost(data *model.SuperlikePost) error
}

func NewDatabase(client DatabaseClient) *Database {
	return &Database{
		Client: client,
	}
}
