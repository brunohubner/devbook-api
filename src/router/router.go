package router

import (
	"github.com/brunohubner/devbook-api/src/router/routes"
	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	return routes.Config(mux.NewRouter())
}
