package middlewares

import (
	"log"
	"net/http"

	"github.com/brunohubner/devbook-api/src/auth"
	"github.com/brunohubner/devbook-api/src/responses"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf(
			"%s %s%s",
			r.Method,
			r.Host,
			r.URL,
		)
		next(w, r)
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateJWT(r); err != nil {
			responses.Error(w, http.StatusUnauthorized, err)
			return
		}

		next(w, r)
	}
}
