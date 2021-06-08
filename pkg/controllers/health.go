package controllers

import (
	"github.com/btassone/crypto/omega/pkg/middleware"
	"github.com/btassone/crypto/omega/pkg/web"
	"net/http"
)

var (
	HealthController = web.Controller{
		Routes: []web.Route{
			{
				Path:      "/api/v1/health",
				Methods:   []string{"GET", "OPTIONS"},
				RouteFunc: GetHealth,
			},
		},
		Middleware: []web.Middleware{
			middleware.CorsMiddleware,
			middleware.TokenMiddleware,
			middleware.LoggingMiddleware,
		},
	}
)

// GetHealth The health route handler
func GetHealth(r *http.Request) (interface{}, error) {
	return map[string]bool{"ok": true}, nil
}
