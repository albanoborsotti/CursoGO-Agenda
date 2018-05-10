package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"ContactAdd",
		"POST",
		"/contactos",
		ContactAdd,
	},
	Route{
		"ContactList",
		"GET",
		"/contactos/",
		ContactList,
	},
	Route{
		"ContactUpdate",
		"PUT",
		"/contactos/{id}",
		ContactUpdate,
	},
	Route{
		"ContactDelete",
		"DELETE",
		"/contactos/{id}",
		ContactDelete,
	},
	Route{
		"ContactSearch",
		"GET",
		"/contactos/{txt}",
		ContactSearch,
	},
}
