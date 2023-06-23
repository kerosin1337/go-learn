package userRequestDto

type CreateUserDto struct {
	Name     interface{} `json:"name" binding:"required"`
	Email    string      `json:"email" binding:"required,email,unique=users;email" example:"test@test.ru"`
	Birthday string      `json:"birthday" binding:"required,date" example:"2015-09-15T14:00:12-00:00"`
	Password string      `json:"password" binding:"required"`
}

type FindAllUserDto struct {
	Name     string `form:"name" binding:"omitempty"`
	Email    string `form:"email" binding:"omitempty"`
	Birthday string `form:"birthday" binding:"omitempty" time_format:"2006-01-02"`
}

type SignInDto struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
