package services

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	AppDb Database
)

type Database struct {
	Conn *gorm.DB
}

type Migrator func(d *Database)

func (d *Database) Connect(migrator Migrator) {
	db, err := gorm.Open(sqlite.Open("./resources/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	AppDb = Database{Conn: db}

	d.Migrate(migrator)
}

func (d *Database) Migrate(migrator func(d *Database)) {
	fmt.Println("Database setup function")

	migrator(d)
}
