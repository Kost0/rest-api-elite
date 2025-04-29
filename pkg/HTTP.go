package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOrders(c *gin.Context) {
	var orders []Order
	db.Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func GetProducts(c *gin.Context) {
	var products []Product
	db.Find(&products)
	c.JSON(http.StatusOK, products)
}

func GetShipments(c *gin.Context) {
	var shipments []Shipment
	db.Find(&shipments)
	c.JSON(http.StatusOK, shipments)
}

func GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	var order Order
	if err := db.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func CreateOrder(c *gin.Context) {
	var newOrder Order
	if err := c.BindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	db.Create(&newOrder)
	c.JSON(http.StatusCreated, newOrder)
}

func UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	var updatedOrder Order
	if err := c.BindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	if err := db.Model(&Order{}).Where("id = ?", id).Updates(updatedOrder).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
		return
	}
	c.JSON(http.StatusOK, updatedOrder)
}

func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&Order{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "order deleted"})
}

func SayHello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello")
}
