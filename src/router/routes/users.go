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
		RequireAuthentication: true,
		Func:                  controllers.FindMany,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodGet,
		RequireAuthentication: true,
		Func:                  controllers.FindOne,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodPut,
		RequireAuthentication: true,
		Func:                  controllers.Update,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodDelete,
		RequireAuthentication: true,
		Func:                  controllers.Delete,
	},
	{
		URI:                   "/users/{userId}/follow",
		Method:                http.MethodPost,
		RequireAuthentication: true,
		Func:                  controllers.FollowUser,
	},
	{
		URI:                   "/users/{userId}/unfollow",
		Method:                http.MethodPost,
		RequireAuthentication: true,
		Func:                  controllers.UnfollowUser,
	},
	{
		URI:                   "/users/{userId}/followers",
		Method:                http.MethodGet,
		RequireAuthentication: true,
		Func:                  controllers.FindFollowers,
	},
	{
		URI:                   "/users/{userId}/following",
		Method:                http.MethodGet,
		RequireAuthentication: true,
		Func:                  controllers.FindFollowing,
	},
	{
		URI:                   "/users/{userId}/password",
		Method:                http.MethodPost,
		RequireAuthentication: true,
		Func:                  controllers.UpdatePassword,
	},
}
