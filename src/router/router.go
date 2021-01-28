package router

import (
	"webapp/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate return an router with all routes
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
