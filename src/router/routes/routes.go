package routes

import (
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

// Route represents all routes of WebApp
type Route struct {
	URI         string
	Method      string
	Handler     func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

// Configure put all routes inside the router.
func Configure(router *mux.Router) *mux.Router {
	routes := routesLogin
	routes = append(routes, routesUsers...)
	routes = append(routes, routeHome)
	routes = append(routes, routesProperties...)

	for _, route := range routes {

		if route.RequireAuth {
			router.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Handler)),
			).Methods(route.Method)
		} else {
			router.HandleFunc(route.URI,
				middlewares.Logger(route.Handler),
			).Methods(route.Method)
		}

	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
