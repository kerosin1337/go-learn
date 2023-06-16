package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string    `json:"name,omitempty" gorm:"<-;embedded"`
	Email    string    `json:"email,omitempty" gorm:"<-;type:varchar(100);unique;embedded"`
	Birthday time.Time `json:"birthday,omitempty" gorm:"<-;embedded"`
}
