package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"monolith-shop/models"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("shop.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
}
