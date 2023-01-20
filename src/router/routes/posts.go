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
		Method:                http.MethodDelete,
		Func:                  controllers.DeletePost,
		RequireAuthentication: true,
	},
	{
		URI:                   "/posts/{postId}/like",
		Method:                http.MethodPost,
		Func:                  controllers.LikeAPost,
		RequireAuthentication: true,
	},
	{
		URI:                   "/posts/{postId}/unlike",
		Method:                http.MethodPost,
		Func:                  controllers.UnlikeAPost,
		RequireAuthentication: true,
	},
}
