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
	if err := database.DB.Where(&model.User{Email: user.Email}).Take(&user).RowsAffected; err > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User does exists"})
		return
	}
	database.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func FindAllUser(c *gin.Context) {
	var users []model.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}
