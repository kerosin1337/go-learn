package common

import (
	"github.com/joho/godotenv"
	"os"
)

func Env(key string) string {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	return os.Getenv(key)
}
