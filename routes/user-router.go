package routes

import (
	userController "example/web-service-gin/controller/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"strings"
)

func ValidationMiddleWare[T any](c *gin.Context) {
	var input T
	var err error
	switch c.Request.Method {
	case "GET":
		err = c.ShouldBindQuery(&input)
		break
	case "POST", "PATCH":
		err = c.ShouldBindJSON(&input)
		break
	}
	if err != nil {
		// Check if it's a validation error
		if errs, ok := err.(validator.ValidationErrors); ok {
			validationErrors := make(map[string]string)

			// Loop through validation errors
			for _, e := range errs {
				fieldName := e.Field()

				// Get the corresponding JSON field name
				if field, ok := reflect.TypeOf(input).FieldByName(fieldName); ok {
					jsonTag := field.Tag.Get("json")
					jsonFieldName := strings.Split(jsonTag, ",")[0]

					// Add the error message to the map
					validationErrors[jsonFieldName] = fmt.Sprintf("Field '%s' validation failed on the rule '%s'", jsonFieldName, e.Tag())
				}
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"error": validationErrors,
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
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
