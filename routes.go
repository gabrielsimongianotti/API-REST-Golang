package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Name(route.Name).
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandleFunc)
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
		"MovieShow",
		"GET",
		"/filme/{id}",
		MovieShow,
	},
	Route{
		"MovieList",
		"GET",
		"/filmes",
		MovieList,
	},
	Route{
		"Contact",
		"GET",
		"/contato",
		Contact,
	},
	Route{
		"MovieAdd",
		"POST",
		"/filme",
		MovieAdd,
	},
	Route{
		"MovieUpdate",
		"PUT",
		"/filme/{id}",
		MovieUpdate,
	},
	Route{
		"MovieRemove",
		"DELETE",
		"/filme/{id}",
		MovieRemove,
	},
}
