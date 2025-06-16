package db

import (
	"github.com/dmvorona/shop/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("shop.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.Cart{}, &models.CartItem{})
}
