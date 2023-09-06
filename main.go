package main

import (
	"net/http"
	"phrase-back/src/db"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

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
