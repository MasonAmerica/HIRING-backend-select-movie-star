package main

import (
	"github.com/gorilla/mux"
)

const QueryVal = "{val}"
const identifier = "X-Identifier"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range AllRoutes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.HandlerFunc).
			Headers(identifier, route.identifier)
	}

	return router
}
