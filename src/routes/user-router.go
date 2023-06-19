package routes

import (
	"example/web-service-gin/src/common/middleware"
	userController "example/web-service-gin/src/module/user"
	userRequestDto "example/web-service-gin/src/module/user/dto/req"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	userRouter := router.Group("user")
	{
		userRouter.GET("/", middleware.ValidationMiddleWare[userRequestDto.FindAllUserDto], userController.FindAllUser)
		userRouter.POST("/", middleware.ValidationMiddleWare[userRequestDto.CreateUserDto], userController.CreateUser)
		userRouter.POST("/sign-in", middleware.ValidationMiddleWare[userRequestDto.SignInDto], userController.SignIn)
		userRouter.GET("/:id", userController.FindOneUser)
	}
}
