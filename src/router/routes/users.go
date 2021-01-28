package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routesUsers = []Route{
	{
		URI:         "/signup",
		Method:      http.MethodGet,
		Handler:     controllers.LoadSignupPage,
		RequireAuth: false,
	},
	{
		URI:         "/users",
		Method:      http.MethodPost,
		Handler:     controllers.CreateUser,
		RequireAuth: false,
	},
}
