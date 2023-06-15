package main

import (
	"example/web-service-gin/database"
	"example/web-service-gin/routes"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.Connect()
	routes.UserRouter(router)
	router.Run("localhost:8080")
}
