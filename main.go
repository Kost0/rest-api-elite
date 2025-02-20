package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Note struct {
	ID          string `json:"id"`
	Customer    string `json:"customer"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Conclusion  string `json:"conclusion"`
	ProductName string `json:"productName"`
	Amount      int    `json:"amount"`
	IsDelivered bool   `json:"isDelivered"`
}

var orders = []Note{
	{ID: "1", Customer: "me", Address: "tam", Phone: "89999999999", Conclusion: "2024-12-01", ProductName: "Air", Amount: 100, IsDelivered: true},
	{ID: "2", Customer: "On", Address: "Tam", Phone: "88005553737", Conclusion: "2024-12-02", ProductName: "Nothing", Amount: 10, IsDelivered: false},
}

func main() {
	router := gin.Default()

	router.GET("/orders", getOrders)

	router.GET("/orders/:id", getOrderByID)

	router.POST("/orders", createOrder)

	router.PUT("/orders/:id", updateOrder)

	router.DELETE("/orders/:id", deleteOrder)

	router.Run(":8080")
}

func getOrders(c *gin.Context) {
	c.JSON(http.StatusOK, orders)
}

func getOrderByID(c *gin.Context) {
	id := c.Param("id")

	for _, order := range orders {
		if order.ID == id {
			c.JSON(http.StatusOK, order)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
}

func createOrder(c *gin.Context) {
	var newOrder Note

	if err := c.BindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	orders = append(orders, newOrder)
	c.JSON(http.StatusCreated, newOrder)
}

func updateOrder(c *gin.Context) {
	id := c.Param("id")
	var updatedOrder Note

	if err := c.BindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	for i, order := range orders {
		if order.ID == id {
			orders[i] = updatedOrder
			c.JSON(http.StatusOK, updatedOrder)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
}

func deleteOrder(c *gin.Context) {
	id := c.Param("id")

	for i, order := range orders {
		if order.ID == id {
			orders = append(orders[:i], orders[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "order deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
}
