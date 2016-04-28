package main

import "net/http"

type Route struct {
	Method      string
	Pattern     string
	Name        string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	Route { "GET", "/",
			"Index", Index,
	},

	Route{ "GET", "/show/{docKey:.+}",
			"Presenter", Presenter,
	},
}
