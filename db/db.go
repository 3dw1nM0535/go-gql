package db

import (
	"github.com/jinzhu/gorm"
	// gorm postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB *gorm.DB
type DB struct {
	*gorm.DB
}

// ConnectDB : connecting DB
func ConnectDB() (*DB, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=demo dbname=rentr password=demo1234 sslmode=disable")

	if err != nil {
		panic(err)
	}

	return &DB{db}, nil
}
