package user

import (
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {

	//	router := router.Group("/users")
	router.POST("/sign", Signin, UserGetByEmail)
	router.POST("/user", UserCreate)
	router.GET("/users", UserGet)
	router.GET("/:username", UserGetByUser)
	router.PUT("/user/:username", UserUpdate)

}
