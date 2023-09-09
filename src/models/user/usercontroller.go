package user

import (
	"os/user"
	"phrase-back/src/db"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func UserGet(c *gin.Context) {
	var users []user.User
	db.DB.Find(&users)
	c.JSON(200, &users)
	return
}
