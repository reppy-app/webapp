package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routesProperties = []Route{
	{
		URI:         "/properties",
		Method:      http.MethodPost,
		Handler:     controllers.CreateProperty,
		RequireAuth: true,
	},
}
