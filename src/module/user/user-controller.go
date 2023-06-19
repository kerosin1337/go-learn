package user

import (
	"example/web-service-gin/src/common/jwt"
	"example/web-service-gin/src/database"
	"example/web-service-gin/src/database/model"
	userRequestDto "example/web-service-gin/src/module/user/dto/req"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateUser godoc
// @Summary Create a new user
// @Description	Create a new user with the provided data
// @Tags Users
// @Accept json
// @Produce	 json
// @Param user body	userRequestDto.CreateUserDto true "User object"
// @Success	201	{object} userResponseDto.UserResponse
// @Failure 422 {object} response.ValidationErrorResponse
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var input = c.MustGet("body").(userRequestDto.CreateUserDto)
	user := model.User{Name: input.Name.(string), Email: input.Email, Birthday: input.Birthday}
	hashPassword, err := user.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hash Error"})
		return
	}
	user.Password = string(hashPassword)
	database.DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func FindAllUser(c *gin.Context) {
	var input = c.MustGet("body").(userRequestDto.FindAllUserDto)
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

func SignIn(c *gin.Context) {
	var input = c.MustGet("body").(userRequestDto.SignInDto)
	var user model.User
	fmt.Print(user.ComparePassword(input.Password))
	if database.DB.Where(model.User{Email: input.Email}).Take(&user).RowsAffected == 0 || user.ComparePassword(input.Password) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Email or Password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user, "token": jwt.GenerateToken(user.ID)})
}

func FindOneUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	if database.DB.Take(&user, id).RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
