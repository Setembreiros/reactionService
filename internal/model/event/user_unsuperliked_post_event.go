package event

type UserUnsuperlikedPostEvent struct {
	Username string `json:"username"`
	PostId   string `json:"postId"`
}

var UserUnsuperlikedPostEventName = "UserUnsuperlikedPostEvent"
