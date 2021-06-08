package web

import (
	"errors"
	"github.com/gorilla/mux"
)

// Controller The controller struct
type Controller struct {
	Routes     []Route
	Middleware []Middleware
}

// SetRoutes set the routes for the current controller
func (c *Controller) SetRoutes(router *mux.Router) {
	// Loop over the kraken exchange routes
	for _, route := range c.Routes {
		// Append default controller middleware
		for _, middleware := range c.Middleware {
			route.Middleware = append(route.Middleware, middleware)
		}

		// Set the route
		route.SetRoute(router)
	}
}

// HasRoutePath does the passed in path exist in the current controller
func (c *Controller) HasRoutePath(route string) bool {
	// Loop over the routes
	for _, r := range c.Routes {
		// If the path is equal to the route passed in return true
		if r.Path == route {
			return true
		}
	}

	// Otherwise false
	return false
}

// GetRoute get the specified route object by path
func (c *Controller) GetRoute(route string) (Route, error) {
	// If the route path exists
	if c.HasRoutePath(route) {
		// Find and return the route
		for _, r := range c.Routes {
			if r.Path == route {
				return r, nil
			}
		}
	}

	// Otherwise generate an error
	return Route{}, errors.New("unable to find matching route")
}
