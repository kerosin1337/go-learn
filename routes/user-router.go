package routes

import (
	"example/web-service-gin/controller/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRouter(router *gin.Engine) {
	userRouter := router.Group("user")
	{
		userRouter.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":  "posted",
				"message": 123,
				"nick":    "123",
			})
		})

		userRouter.POST("", user.CreateUser)
	}
}
