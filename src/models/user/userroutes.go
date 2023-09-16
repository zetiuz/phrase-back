package user

import (
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {

	routes := router.Group("/users")
	routes.GET("", UserGet)

}
