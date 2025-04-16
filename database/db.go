package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Строка подключения для PostgreSQL
	dsn := "user=postgres password=1111 dbname=phonego port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	// Автоматическое создание таблиц
	DB = db
}
