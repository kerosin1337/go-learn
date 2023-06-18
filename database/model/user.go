package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string    `json:"name,omitempty" gorm:"<-;embedded"`
	Email    string    `json:"email,omitempty" gorm:"<-;type:varchar(100);unique;embedded"`
	Birthday time.Time `json:"birthday,omitempty" gorm:"<-;embedded"`
	Password string    `json:"-" gorm:"<-;embedded;not null"`
}

func (user *User) HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (user *User) ComparePassword(password string) error {
	hashPass, err := user.HashPassword(password)
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword([]byte(user.Password), hashPass)
}
