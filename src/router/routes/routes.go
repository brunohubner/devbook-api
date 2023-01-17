package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                   string
	Method                string
	RequireAuthentication bool
	Func                  func(http.ResponseWriter, *http.Request)
}

func Config(r *mux.Router) *mux.Router {
	for _, route := range userRoutes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	return r
}
