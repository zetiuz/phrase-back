package user

import (
	"phrase-back/src/db"
	"time"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Username    string    `json:"username"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Birthdate   time.Time `json:"birthdate"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Language_id int64     `json:"language"`
	Status      bool      `json:"status"`
}

func UserGet(c *gin.Context) {
	var user []Users
	db.DB.Table("users").Select("*").Scan(&user)
	c.JSON(200, &user)
	return
}
