package comments

import "github.com/gin-gonic/gin"

func CommentsRouter(router *gin.Engine) {

	router.POST("/comments", CommentsCreate)
}
