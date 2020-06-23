package main

import (
	"github.com/mschlech/go-template/cmd/handler"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"healthCheck",
		"GET",
		"/healthcheck/",
		handler.GetHealthCheck,
	},
}
