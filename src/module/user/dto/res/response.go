package userResponseDto

import "time"

type UserResponse struct {
	ID        uint      `json:"ID" example:"36" binding:"required"`
	CreatedAt string    `json:"CreatedAt" example:"2023-06-19T12:23:08.904269+07:00" binding:"required"`
	UpdatedAt string    `json:"UpdatedAt" example:"2023-06-19T12:23:08.904269+07:00" binding:"required"`
	DeletedAt string    `json:"DeletedAt" example:"null"`
	Name      string    `json:"name" example:"test" binding:"required"`
	Email     string    `json:"email" example:"test@test.ru" binding:"required"`
	Birthday  time.Time `json:"birthday" example:"2015-09-15T14:00:12Z" binding:"required"`
}
