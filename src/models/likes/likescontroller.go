package likes

import (
	"phrase-back/src/db"

	"github.com/gin-gonic/gin"
)

type LikesRequest struct {
	User_id string `json:"user_id"`
	Post_id int64  `json:"username"`
}

func LikesCreate(c *gin.Context) {
	body := LikesRequest{}

	c.BindJSON(&body)

	create := &Likes{User_id: body.User_id, Post_id: body.Post_id}

	result := db.DB.Create(&create)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(200, &create)
}
func LikesDelete(c *gin.Context) {

	user := c.Param("user_id")
	var likes Likes

	db.DB.Where("post_id= ?", &likes.Post_id).Delete(&likes, user)
	c.JSON(200, gin.H{"deleted": true})
	return
}
