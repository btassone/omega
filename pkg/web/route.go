package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// RouteFunc The logic to run that generates the data to return
type RouteFunc func(r *http.Request) (interface{}, error)

// HttpHandlerFunc The normal handler that returns for the route with default output logic
type HttpHandlerFunc func(w http.ResponseWriter, r *http.Request)

// Route An individual route used for path construction
type Route struct {
	Path       string
	RouteFunc  RouteFunc
	Methods    []string
	Middleware []Middleware
}

// SetRoute Sets the route on the router with the handler func and methods
func (route *Route) SetRoute(router *mux.Router) {
	configuredRouteFunc := RegisterMiddleware(route.WrapHandlerFunc(route.RouteFunc), route.Middleware)

	router.HandleFunc(route.Path, configuredRouteFunc).Methods(route.Methods...)
}

// WrapHandlerFunc Wraps the route logic handler with the http handler the server expects
func (route *Route) WrapHandlerFunc(handlerFunc RouteFunc) func(w http.ResponseWriter, r *http.Request) {
	// Return the expected function required for the route handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Call the RouteFunc and get the data
		data, err := handlerFunc(r)
		// Setup the json encoder
		encoder := json.NewEncoder(w)

		// If there is an error print to console and return error
		if err != nil {
			log.Println(err)
			_ = encoder.Encode(map[string]string{
				"error": fmt.Sprintf("%s", err),
			})
			// Otherwise encode the data
		} else {
			_ = encoder.Encode(data)
		}
	}
}
