package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateConnection() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("DB_PASSWORD")

	config := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, dbname, pass)

	return gorm.Open("postgres", config)
}
