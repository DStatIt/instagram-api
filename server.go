package IGapi

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
		"Test",
		"GET",
		"/test/{test}",
		TestHandler,
	},
}

func init() {
	router := NewRouter()
}
