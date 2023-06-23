package main

import (
	_ "example/web-service-gin/docs"
	"example/web-service-gin/src/common/validate"
	"example/web-service-gin/src/database"
	"example/web-service-gin/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "net/http"
)

// @title Your Gin API
// @version 0.2
// @description	This is a sample Gin API with Swagger documentation
// @host localhost:12300
// @BasePath /
func main() {
	router := gin.Default()
	database.Connect()
	routes.UserRouter(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// TODO: Custom Validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("unique", validate.UniqueValidate)
		v.RegisterValidation("date", validate.DateValidate)
	}
	router.Run("localhost:12300")
}
