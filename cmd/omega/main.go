package main

import (
	"github.com/btassone/crypto/omega/pkg/services"
)

// main The main function of the application
func main() {
	// Start the api web server
	services.ApiServer.Start()
}
