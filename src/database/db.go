package database

import (
	"example/web-service-gin/src/common"
	"example/web-service-gin/src/database/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

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
	db.AutoMigrate(&model.User{})

	DB = db
}
