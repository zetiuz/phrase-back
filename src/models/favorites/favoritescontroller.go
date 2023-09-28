package favorites

import (
	"phrase-back/src/db"

	"github.com/gin-gonic/gin"
)

type FavoritesRequest struct {
	User_id string `json:"user_id"`
	Post_id int64  `json:"post_id"`
}

func FavoritesCreate(c *gin.Context) {
	body := FavoritesRequest{}

	c.BindJSON(&body)

	create := &Favorites{User_id: body.User_id, Post_id: body.Post_id}

	result := db.DB.Create(&create)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(200, &create)
}
