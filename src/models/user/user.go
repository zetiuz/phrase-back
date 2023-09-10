package user

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username    string
	Name        string
	Description string
	Birthdate   time.Time
	Email       string
	Password    string
	Language_id int64
	Status      bool
}

//
