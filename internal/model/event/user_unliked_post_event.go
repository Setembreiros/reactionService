package event

type UserUnlikedPostEvent struct {
	Username string `json:"username"`
	PostId   string `json:"postId"`
}

var UserUnlikedPostEventName = "UserUnlikedPostEvent"
