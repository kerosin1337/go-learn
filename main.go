package main

import (
	"example/web-service-gin/src/database"
	"example/web-service-gin/src/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "net/http"
)

// @Summary Get user by ID
// @Description Get a user by their ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Router /users/{id} [get]
func getUserByID(c *gin.Context) {
	// Handle get user by ID
}

type User struct {
	// ID this is userid
	ID   int    `json:"id"`
	Name string `json:"name" example:"asd"`
}

func main() {
	router := gin.Default()
	database.Connect()
	routes.UserRouter(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/users/:id", getUserByID)
	router.Run("localhost:8080")
}
