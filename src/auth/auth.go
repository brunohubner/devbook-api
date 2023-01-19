package auth

import (
	"time"

	"github.com/brunohubner/devbook-api/src/config"
	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJWT(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6)
	permissions["userId"] = userID
	permissions["authorized"] = true

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString(config.JwtSecret)
}
