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
	{
		URI:         "/properties/{{propertyId}}/edit",
		Method:      http.MethodGet,
		Handler:     controllers.LoadEditPropertyPage,
		RequireAuth: true,
	},
}
