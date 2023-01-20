package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

func ValidateJWT(r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, getJwtSecret)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return errors.New("Invalid token")
	}

	return nil
}

func ExtractUserID(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, getJwtSecret)
	if err != nil {
		return 0, err
	}

	permissions, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("Invalid token")
	}

	userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func getJwtSecret(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signature method: %v", token.Header["alg"])
	}

	return config.JwtSecret, nil
}
