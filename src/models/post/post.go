package post

import (
	"time"
)

type Posts struct {
	Id          string
	User_id     string
	Image       string
	Body        string
	CreatedAt   time.Time
	Thematic_id int64
}

func (Posts) TableName() string {
	return "posts"
}
