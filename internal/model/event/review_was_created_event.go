package event

type ReviewWasCreatedEvent struct {
	ReviewId  uint64 `json:"reviewId"`
	Username  string `json:"username"`
	PostId    string `json:"postId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Rating    int    `json:"rating"`
	CreatedAt string `json:"createdAt"`
}

var ReviewWasCreatedEventName = "ReviewWasCreatedEvent"
