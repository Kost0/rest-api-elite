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
		protected.GET("/orders", getOrders)
		protected.POST("/orders", createOrder)
	}

	router.GET("/orders/:id", getOrderByID)

	router.PUT("/orders/:id", updateOrder)

	router.DELETE("/orders/:id", deleteOrder)

	router.Run(":8080")
}
