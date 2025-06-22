package db

import (
	"fmt"
	"os"

	"github.com/dmvorona/shop/models"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	server := os.Getenv("AZURE_SQL_SERVER")
	database := os.Getenv("AZURE_SQL_DATABASE")
	username := os.Getenv("AZURE_SQL_USER")
	password := os.Getenv("AZURE_SQL_PASSWORD")

	dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;",
		server, username, password, database)

	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to Azure SQL Database: " + err.Error())
	}

	err = DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.Cart{}, &models.CartItem{})
	if err != nil {
		panic("Failed to migrate database schema: " + err.Error())
	}
}
