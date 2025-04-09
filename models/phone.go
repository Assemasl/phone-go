package models

import "gorm.io/gorm"

// Phone структура для модели телефона
type Phone struct {
	gorm.Model          // Включает ID, CreatedAt, UpdatedAt, DeletedAt
	Brand       string  `json:"brand"`        // Марка телефона
	ModelName   string  `json:"model_name"`   // Модель телефона (переименовано с "Model" для избежания конфликта)
	Price       float64 `json:"price"`        // Цена телефона
	ReleaseDate string  `json:"release_date"` // Дата выпуска
}
