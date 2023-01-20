package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/brunohubner/devbook-api/src/config"
	"github.com/brunohubner/devbook-api/src/router"
)

func main() {
	config.LoadEnvVars()

	router := router.Generate()

	host := fmt.Sprintf(":%d", config.Port)

	fmt.Printf("API listening on http://localhost:%d\n", config.Port)

	log.Fatal(http.ListenAndServe(host, router))
}
