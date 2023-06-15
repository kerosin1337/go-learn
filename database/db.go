package database

import (
	"example/web-service-gin/common"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}
type Author struct {
	gorm.Model
	Name  string
	Email string
}

type Blog struct {
	gorm.Model
	ID      int
	Author  Author `gorm:"embedded"`
	Upvotes int32
	Text    string `gorm:"type:text"`
}

func Connect() {
	user := common.Env("PG_USER")
	pass := common.Env("PG_PASS")
	host := common.Env("PG_HOST")
	port := common.Env("PG_PORT")
	name := common.Env("PG_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, name, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Author{})
	db.AutoMigrate(&Blog{})
}
