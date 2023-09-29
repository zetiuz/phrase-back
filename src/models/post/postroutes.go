package post

import (
	"github.com/gin-gonic/gin"
)

func PostRouter(router *gin.Engine) {

	router.POST("/post", PostCreate)
	router.GET("/post", PostGet)
	router.PUT("/post/:id", PostUpdate)
	router.DELETE("/post/:id", PostDelete)

}
