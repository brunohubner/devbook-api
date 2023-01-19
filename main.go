package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/brunohubner/devbook-api/src/config"
	"github.com/brunohubner/devbook-api/src/router"
)

// func init() {
// 	jwtSecret := make([]byte, 64)

// 	if _, err := rand.Read(jwtSecret); err != nil {
// 		log.Fatal(err)
// 	}

// 	stringBase64 := base64.StdEncoding.EncodeToString(jwtSecret)
// 	fmt.Println(stringBase64)
// }

func main() {
	config.LoadEnvVars()

	router := router.Generate()

	host := fmt.Sprintf(":%d", config.Port)

	fmt.Printf("API listening on http://localhost:%d", config.Port)

	log.Fatal(http.ListenAndServe(host, router))
}
