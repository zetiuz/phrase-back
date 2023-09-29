package comments

import (
	"phrase-back/src/db"
	"time"

	"github.com/gin-gonic/gin"
)

type CommentsRequest struct {
	Post_id   int64     `json:"post_id"`
	User_id   string    `json:"user_id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
}

func CommentsCreate(c *gin.Context) {
	body := CommentsRequest{}

	c.ShouldBindJSON(&body)

	create := &Comments{Post_id: body.Post_id, User_id: body.User_id, Body: body.Body, CreatedAt: time.Now()}

	result := db.DB.Create(&create)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(200, &create)
}
