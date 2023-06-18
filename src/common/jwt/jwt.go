package jwt

import (
	"example/web-service-gin/src/common"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

func GenerateToken(id uint) string {
	tokenLife, err := strconv.Atoi(common.Env("TOKEN_LIFE"))
	if err != nil {
		tokenLife = 1
	}
	tokenSecret := common.Env("TOKEN_SECRET")
	if tokenSecret == "" {
		tokenSecret = "test"
	}
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLife)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	res, _ := token.SignedString([]byte(tokenSecret))
	return res
}
