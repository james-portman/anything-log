package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"LogIndex",
		"GET",
		"/logs",
		LogIndex,
	},
	Route{
		"LogCreate",
		"POST",
		"/logs",
		LogCreate,
	},
	Route{
		"LogShow",
		"GET",
		"/logs/{logId}",
		LogShow,
	},
}
