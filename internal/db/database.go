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
	GetLikePost(postId, username string) (*model.LikePost, error)
	DeleteLikePost(data *model.LikePost) error
}

func NewDatabase(client DatabaseClient) *Database {
	return &Database{
		Client: client,
	}
}
