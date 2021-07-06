package main

import (
	"github.com/btassone/crypto/omega/cmd/omega/controllers"
	"github.com/btassone/crypto/omega/pkg/services"
)

// main The main function of the application
func main() {
	// Start the database connection
	services.AppDb.Connect(func(d *services.Database) {
		// TODO: Test code, please remove
		d.Conn.AutoMigrate(controllers.Foo{})
		d.Conn.Create(&controllers.Foo{Name: "test", Description: "This is a test description"})
	})

	// Add test controller
	services.ApiServer.AddController(controllers.FooController)

	// Start the api web server
	services.ApiServer.Start()
}
