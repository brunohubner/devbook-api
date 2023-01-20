package routes

import (
	"net/http"

	"github.com/brunohubner/devbook-api/src/controllers"
)

var postsRoutes = []Route{
	{
		URI:                   "/posts",
		Method:                http.MethodPost,
		Func:                  controllers.CreatePost,
		RequireAuthentication: true,
	},
	{
		URI:                   "/posts",
		Method:                http.MethodGet,
		Func:                  controllers.FindPosts,
		RequireAuthentication: true,
	},
	{
		URI:                   "/posts/{postId}",
		Method:                http.MethodGet,
		Func:                  controllers.FindPostById,
		RequireAuthentication: true,
	},
	{
		URI:                   "/posts/{postId}",
		Method:                http.MethodPut,
		Func:                  controllers.UpdatePost,
		RequireAuthentication: true,
	},
	{
		URI:                   "/posts/{postId}",
		Method:                http.MethodPost,
		Func:                  controllers.DeletePost,
		RequireAuthentication: true,
	},
}
