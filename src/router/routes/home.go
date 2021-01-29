package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routeHome = Route{
	URI:         "/home",
	Method:      http.MethodGet,
	Handler:     controllers.LoadHome,
	RequireAuth: true,
}
