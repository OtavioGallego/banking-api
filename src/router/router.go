package router

import (
	"github.com/OtavioGallego/banking-api/src/router/routes"
	"github.com/gorilla/mux"
)

// Generate return the router with all routes configured
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
