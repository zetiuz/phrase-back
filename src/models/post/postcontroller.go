package post

import (
	"time"

	"phrase-back/src/db"

	"github.com/gin-gonic/gin"
)

type PostRequest struct {
	Id          string    `json:"id"`
	User_id     string    `json:"user_id"`
	Imagen      string    `json:"imagen"`
	Body        string    `json:"body"`
	CreatedAt   time.Time `json:"createdAt"`
	Thematic_id int64     `json:"thematic_id"`
}

func PostCreate(c *gin.Context) {
	body := PostRequest{}

	c.BindJSON(&body)

	create := &Posts{Id: body.Id, User_id: body.User_id, Imagen: body.Imagen, Body: body.Body, CreatedAt: body.CreatedAt, Thematic_id: body.Thematic_id}

	result := db.DB.Create(&create)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(200, &create)
}
func PostUpdate(c *gin.Context) {

	id := c.Param("id")
	var post Posts
	db.DB.First(&post, "id", id)

	body := PostRequest{}
	c.BindJSON(&body)
	data := &Posts{Imagen: body.Imagen, Body: body.Body}

	result := db.DB.Model(&post).Where("id", id).Updates(data)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, &post)
}
