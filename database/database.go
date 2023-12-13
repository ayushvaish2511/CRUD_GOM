package database

import (
	// "fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB 

func Connect() error {
	dsn := "host=localhost user=postgres password=admin dbname=crud port=5432 sslmode=disable"
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db
	return nil
}
  