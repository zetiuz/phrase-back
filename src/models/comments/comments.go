package comments

import "time"

type Comments struct {
	Post_id   int64
	User_id   string
	Body      string
	CreatedAt time.Time
}

func (Comments) TableName() string {
	return "comments"
}
