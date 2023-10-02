package post

import (
	"github.com/gin-gonic/gin"
)

func PostRouter(router *gin.Engine) {

	router.POST("/post", TokenAuthMiddleware(), PostCreate)
	router.GET("/post", PostGet)
	router.PUT("/post/:id", PostUpdate)
	router.DELETE("/post/:id", PostDelete)
	router.GET("/post/:user_id", PostGetByUsername)
	router.GET("/posts/:thematic_id", PostGetByThematics)
	router.GET("/searchposts/:word", PostGetByWord)

}
