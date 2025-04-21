package event

type UserLikedPostEvent struct {
	Username string `json:"username"`
	PostId   string `json:"postId"`
}

var UserLikedPostEventName = "UserLikedPostEvent"
