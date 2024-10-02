package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Подключение к базе данных
func connectDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(getDatabaseURL()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
