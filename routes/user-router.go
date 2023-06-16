package routes

import (
	"example/web-service-gin/controller/user"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	userRouter := router.Group("user")
	{
		userRouter.GET("/", user.FindAllUser)

		userRouter.POST("", user.CreateUser)
	}
}
