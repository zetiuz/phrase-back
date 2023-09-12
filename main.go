package main

import (
	"net/http"
	"phrase-back/src/db"
	"phrase-back/src/models/user"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go.",
		})
	})
	router.POST("", user.UserCreate)
	router.GET("/users", user.UserGet)
	router.GET("/:username", user.UserGetByUser)
	router.PUT("/:username", user.UserUpdate)
	router.Run()
}
func init() {
	db.ConnectToDB()
}
