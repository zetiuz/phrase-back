package user

import (
	"phrase-back/src/db"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Language_id int64  `json:"language"`
	Status      bool   `json:"status"`
}

func UserCreate(c *gin.Context) {

	body := UserRequest{}

	c.BindJSON(&body)

	create := &Users{Username: body.Username, Name: body.Name, Description: body.Description, Email: body.Email, Password: body.Password, Language_id: body.Language_id, Status: body.Status}

	result := db.DB.Create(&create)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(200, &create)
}
func UserGet(c *gin.Context) {
	var user []Users
	db.DB.Table("user").Select("*").Scan(&user)
	c.JSON(200, &user)
	return
}

func UserGetByUser(c *gin.Context) {
	Username := c.Param("username")
	var user Users
	db.DB.First(&user, "username", Username)
	c.JSON(200, &user)
	return
}

func UserUpdate(c *gin.Context) {

	username := c.Param("username")
	var user Users
	db.DB.First(&user, username)

	body := UserRequest{}
	c.BindJSON(&body)
	data := &Users{Name: body.Name, Description: body.Description, Email: body.Email, Password: body.Password}

	result := db.DB.Model(&user).Updates(data)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, &user)
}
