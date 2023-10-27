package jwt

import (
	"fmt"
	"github.com/dashixiong47/KK_BBS/config"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var mySigningKey = []byte(config.SettingsConfig.Application.SigningKey)

func CreateToken(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = user.Username
	claims["id"] = user.ID
	timeout := 0
	if config.SettingsConfig.Application.Mode == "dev" {
		timeout = 999999
	} else {
		timeout = config.SettingsConfig.Application.TokenTimeout
	}

	claims["exp"] = time.Now().Add(time.Minute * time.Duration(timeout)).Unix()
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})
}
