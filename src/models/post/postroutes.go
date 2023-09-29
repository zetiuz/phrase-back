package post

import (
	"github.com/gin-gonic/gin"
)

func PostRouter(router *gin.Engine) {

	router.POST("/post", PostCreate)

}
