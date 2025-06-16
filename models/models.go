package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Orders   []Order
}

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	Stock       int
}

type Order struct {
	gorm.Model
	UserID    uint
	ProductID uint
	Quantity  int
}

type Cart struct {
	gorm.Model
	UserID    uint
	CartItems []CartItem
}

type CartItem struct {
	gorm.Model
	CartID    uint
	ProductID uint
	Quantity  int

	Product Product
}
