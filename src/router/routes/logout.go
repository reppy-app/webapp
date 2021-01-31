package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routeLogout = Route{
	URI:         "/logout",
	Method:      http.MethodGet,
	Handler:     controllers.Logout,
	RequireAuth: true,
}
