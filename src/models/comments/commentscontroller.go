package comments

import "time"

type CommentsRequest struct {
	Post_id   int64     `json:"post_id"`
	User_id   string    `json:"user_id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
}
