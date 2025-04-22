package event

type UserSuperlikedPostEvent struct {
	Username string `json:"username"`
	PostId   string `json:"postId"`
}

var UserSuperlikedPostEventName = "UserSuperlikedPostEvent"
