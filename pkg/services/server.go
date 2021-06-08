package services

import (
	"github.com/btassone/crypto/omega/pkg/controllers"
	"github.com/btassone/crypto/omega/pkg/web"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var (
	ApiServer = Server{
		Address:      "omega.local:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Controllers: []web.Controller{
			controllers.HealthController,
		},
	}
)

// Server The controllers server struct
type Server struct {
	Address      string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	Controllers  []web.Controller
	Handler      http.Handler
	Router       *mux.Router
	Instance     *http.Server
}

// Start Starts the controllers server
func (s *Server) Start() {
	// Set the router
	s.SetRouter()

	// Set the routes
	s.SetRoutes()

	// Launch the server
	s.Instance = &http.Server{
		Handler:      s.Router,
		Addr:         s.Address,
		WriteTimeout: s.WriteTimeout,
		ReadTimeout:  s.ReadTimeout,
	}

	// Listen and serve
	log.Fatal(s.Instance.ListenAndServe())
}

// SetRouter Set the instance of the router on the server
func (s *Server) SetRouter() {
	// App router
	s.Router = mux.NewRouter()
}

// SetRoutes Set the server routes and their associated handlers through the controllers
func (s *Server) SetRoutes() {
	// Loop through the controllers
	for _, controller := range s.Controllers {
		// Set the controller routes for the controller on the router
		controller.SetRoutes(s.Router)
	}
}
