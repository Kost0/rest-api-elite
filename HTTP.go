package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getOrders(c *gin.Context) {
	var orders []Order
	db.Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func getProducts(c *gin.Context) {
	var products []Product
	db.Find(&products)
	c.JSON(http.StatusOK, products)
}

func getShipments(c *gin.Context) {
	var shipments []Shipment
	db.Find(&shipments)
	c.JSON(http.StatusOK, shipments)
}

func getOrderByID(c *gin.Context) {
	id := c.Param("id")
	var book Order
	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func createOrder(c *gin.Context) {
	var newBook Order
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	db.Create(&newBook)
	c.JSON(http.StatusCreated, newBook)
}

func updateOrder(c *gin.Context) {
	id := c.Param("id")
	var updatedBook Order
	if err := c.BindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	if err := db.Model(&Order{}).Where("id = ?", id).Updates(updatedBook).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
		return
	}
	c.JSON(http.StatusOK, updatedBook)
}

func deleteOrder(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&Order{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "order deleted"})
}
