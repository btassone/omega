package controllers

import (
	"github.com/btassone/crypto/omega/pkg/middleware"
	"github.com/btassone/crypto/omega/pkg/services"
	"github.com/btassone/crypto/omega/pkg/web"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

type Foo struct {
	gorm.Model
	Name        string
	Description string
}

var (
	FooController = web.Controller{
		Name: "FooController",
		Routes: []web.Route{
			{
				Path:      "/api/v1/foo/{name:[a-zA-Z-]+}",
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
	name := mux.Vars(r)["name"]

	var foo Foo
	services.AppDb.Conn.First(&foo, "name = ?", name)

	return foo, nil
}
