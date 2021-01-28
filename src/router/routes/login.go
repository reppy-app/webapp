package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routesLogin = []Route{
	{
		URI:         "/",
		Method:      http.MethodGet,
		Handler:     controllers.LoadLoginPage,
		RequireAuth: false,
	},
	{
		URI:         "/login",
		Method:      http.MethodGet,
		Handler:     controllers.LoadLoginPage,
		RequireAuth: false,
	},
}
