package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	initDB()

	router := gin.Default()

	router.POST("/login", login)

	protected := router.Group("/")
	protected.Use(authMiddleware())
	{
		protected.PUT("/orders/:id", updateOrder)
		protected.DELETE("/orders/:id", deleteOrder)
		protected.POST("/orders", createOrder)
	}

	router.GET("/orders", getOrders)

	router.GET("/products", getProducts)

	router.GET("/shipments", getShipments)

	router.GET("/orders/:id", getOrderByID)

	router.Run(":8080")
}
