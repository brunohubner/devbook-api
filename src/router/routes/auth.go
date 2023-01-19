package routes

import (
	"net/http"

	"github.com/brunohubner/devbook-api/src/controllers"
)

var authRoutes = []Route{{
	URI:                   "/login",
	Method:                http.MethodPost,
	Func:                  controllers.Login,
	RequireAuthentication: false,
}}
