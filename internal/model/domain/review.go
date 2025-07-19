package model

import "time"

type Review struct {
	Id        uint64    `json:"id"`
	Username  string    `json:"username"`
	PostId    string    `json:"postId"`
	Content   string    `json:"content"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
