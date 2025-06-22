package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	IsAdmin  bool

	Cart   Cart
	Orders []Order
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	Stock       int

	CartItems  []CartItem
	OrderItems []OrderItem
}

type Cart struct {
	gorm.Model
	UserID    uint `gorm:"uniqueIndex"` // 1:1 relation with User
	CartItems []CartItem
}

type CartItem struct {
	gorm.Model
	CartID    uint
	ProductID uint
	Quantity  int

	Product Product
}

type Order struct {
	gorm.Model
	UserID      uint
	Status      string // e.g., "pending", "shipped", etc.
	TotalAmount float64
	OrderDate   int64 `gorm:"autoCreateTime:milli"` // or use time.Time
	OrderItems  []OrderItem
}

type OrderItem struct {
	gorm.Model
	OrderID   uint
	ProductID uint
	Quantity  int
	Price     float64

	Product Product
}
