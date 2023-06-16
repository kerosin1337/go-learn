package user

import (
	"example/web-service-gin/database"
	"example/web-service-gin/database/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CreateUserDto struct {
	Name     string    `json:"name" binding:"required"`
	Email    string    `json:"email" binding:"required,email"`
	Birthday time.Time `json:"birthday" binding:"required" time_format:"2006-01-02"`
}

func CreateUser(c *gin.Context) {
	var input CreateUserDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := model.User{Name: input.Name, Email: input.Email, Birthday: input.Birthday}
	database.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}
