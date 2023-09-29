package main

import (
	"net/http"
	"phrase-back/src/db"
	"phrase-back/src/models/comments"
	"phrase-back/src/models/post"
	"phrase-back/src/models/user"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	user.UserRouter(router)
	post.PostRouter(router)
	comments.CommentsRouter(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go.",
		})
	})

	router.Run()
}
func init() {
	db.ConnectToDB()
}
