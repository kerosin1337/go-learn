package routes

import (
	"example/web-service-gin/common/middleware"
	userController "example/web-service-gin/controller/user"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	userRouter := router.Group("user")
	{
		userRouter.GET("/", middleware.ValidationMiddleWare[userController.FindAllUserDto], userController.FindAllUser)
		userRouter.POST("/", middleware.ValidationMiddleWare[userController.CreateUserDto], userController.CreateUser)
		userRouter.POST("/sign-in", middleware.ValidationMiddleWare[userController.SignInDto], userController.SignIn)
	}
}
