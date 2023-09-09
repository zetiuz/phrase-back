package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string
	Name        string
	Description string
	//	Birthdate   time.Time
	//	Email       string
	//	Password    string
	//	Language_id int64
	//	Status      bool
}

func (User) TableName() string {
	return "users"
}
