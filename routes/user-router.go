package routes

import (
	userController "example/web-service-gin/controller/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidationMiddleWare[T any](c *gin.Context) {
	var input T
	switch c.Request.Method {
	case "GET":
		if err := c.BindQuery(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
		}
		break
	case "POST":
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
		}
		break
	}

	c.Set("body", input)
	c.Next()
}
func UserRouter(router *gin.Engine) {
	userRouter := router.Group("user")
	{
		userRouter.GET("/", ValidationMiddleWare[userController.FindAllUserDto], userController.FindAllUser)

		userRouter.POST("/", ValidationMiddleWare[userController.CreateUserDto], userController.CreateUser)
	}
}
