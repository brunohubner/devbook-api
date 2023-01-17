package routes

import (
	"net/http"

	"github.com/brunohubner/devbook-api/src/controllers"
)

var userRoutes = []Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		RequireAuthentication: false,
		Func:                  controllers.Create,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		RequireAuthentication: false,
		Func:                  controllers.FindAll,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodGet,
		RequireAuthentication: false,
		Func:                  controllers.FindOne,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodPut,
		RequireAuthentication: false,
		Func:                  controllers.Update,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodDelete,
		RequireAuthentication: false,
		Func:                  controllers.Delete,
	},
}
