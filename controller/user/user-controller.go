package user

import (
	"example/web-service-gin/database"
	"example/web-service-gin/database/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CreateUserDto struct {
	Name     interface{} `json:"name" binding:"required,alphanum"`
	Email    string      `json:"email" binding:"required,email"`
	Birthday time.Time   `json:"birthday" binding:"required" time_format:"2006-01-02"`
}

type FindAllUserDto struct {
	Name     string `form:"name" binding:"omitempty"`
	Email    string `form:"email" binding:"omitempty,email"`
	Birthday string `form:"birthday" binding:"omitempty" time_format:"2006-01-02"`
}

func CreateUser(c *gin.Context) {
	var input = c.MustGet("body").(CreateUserDto)
	user := model.User{Name: input.Name.(string), Email: input.Email, Birthday: input.Birthday}
	if err := database.DB.Where(model.User{Email: user.Email}).Take(&user).RowsAffected; err > 0 {
		fmt.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "User does exists"})
		return
	}
	database.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func FindAllUser(c *gin.Context) {
	var input = c.MustGet("body").(FindAllUserDto)
	var users []model.User
	var query = database.DB.Model(&model.User{})
	if input.Name != "" {
		query.Where("name ILIKE ?", "%"+input.Name+"%")
	}
	if input.Email != "" {
		query.Where("email = ?", input.Email)
	}
	if input.Birthday != "" {
		query.Where("birthday = ?", input.Birthday)
	}
	query.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}