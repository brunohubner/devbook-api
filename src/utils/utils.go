package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"strings"
)

func StandardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func GenerateJwtSecret() string {
	jwtSecret := make([]byte, 64)

	if _, err := rand.Read(jwtSecret); err != nil {
		log.Fatal(err)
	}

	return base64.StdEncoding.EncodeToString(jwtSecret)
}
