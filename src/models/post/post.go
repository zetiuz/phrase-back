package post

import (
	"time"
)

type Posts struct {
	Id        string
	User_id   string
	Imagen    string
	Body      string
	CreatedAt time.Time
}
