package controllers

import (
	"github.com/btassone/crypto/omega/pkg/middleware"
	"github.com/btassone/crypto/omega/pkg/web"
	"net/http"
)

var (
	FooController = web.Controller{
		Name: "FooController",
		Routes: []web.Route{
			{
				Path:      "/api/v1/foo",
				Methods:   []string{"GET", "OPTIONS"},
				RouteFunc: GetFoo,
			},
		},
		Middleware: []web.Middleware{
			middleware.CorsMiddleware,
			middleware.TokenMiddleware,
			middleware.LoggingMiddleware,
		},
	}
)

// GetFoo The health route handler
func GetFoo(r *http.Request) (interface{}, error) {
	return map[string]bool{"ok": true}, nil
}
