package routes

import (
	"net/http"

	"github.com/brunohubner/devbook-api/src/middlewares"
	"github.com/gorilla/mux"
)

type Route struct {
	URI                   string
	Method                string
	RequireAuthentication bool
	Func                  func(http.ResponseWriter, *http.Request)
}

func Config(r *mux.Router) *mux.Router {
	includeRoutes(r, userRoutes)
	includeRoutes(r, authRoutes)
	includeRoutes(r, postsRoutes)

	return r
}

func includeRoutes(r *mux.Router, routes []Route) {
	for _, route := range routes {
		if route.RequireAuthentication {
			r.HandleFunc(
				route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Func)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(
				route.URI,
				middlewares.Logger(route.Func),
			).Methods(route.Method)
		}
	}
}
