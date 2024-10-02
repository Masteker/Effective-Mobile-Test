package main

import "gorm.io/gorm"

// Модель для песни
type Song struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Group       string `json:"group"`
	Title       string `json:"song"`
	ReleaseDate string `json:"release_date,omitempty"`
	Text        string `json:"text,omitempty"`
	Link        string `json:"link,omitempty"`
}

// Функция для миграции базы данных
func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&Song{})
}
