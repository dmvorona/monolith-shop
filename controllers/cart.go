package controllers

import (
	"github.com/dmvorona/shop/db"
	"github.com/dmvorona/shop/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddToCartInput struct {
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

func AddToCart(c *gin.Context) {
	var input AddToCartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if cart exists or create
	var cart models.Cart
	db.DB.Where("user_id = ?", input.UserID).FirstOrCreate(&cart, models.Cart{UserID: input.UserID})

	// Check if item already in cart
	var cartItem models.CartItem
	result := db.DB.Where("cart_id = ? AND product_id = ?", cart.ID, input.ProductID).First(&cartItem)

	if result.RowsAffected > 0 {
		cartItem.Quantity += input.Quantity
		db.DB.Save(&cartItem)
	} else {
		cartItem = models.CartItem{
			CartID:    cart.ID,
			ProductID: input.ProductID,
			Quantity:  input.Quantity,
		}
		db.DB.Create(&cartItem)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Added to cart"})
}

func ViewCart(c *gin.Context) {
	userID := c.Query("user_id")
	var cart models.Cart
	if err := db.DB.Preload("CartItems.Product").Where("user_id = ?", userID).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	c.JSON(http.StatusOK, cart)
}

func RemoveFromCart(c *gin.Context) {
	cartItemID := c.Param("id")
	db.DB.Delete(&models.CartItem{}, cartItemID)
	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}
