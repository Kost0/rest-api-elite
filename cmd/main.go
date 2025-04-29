package main

import (
	"github.com/gin-gonic/gin"
	"rest_api1/pkg"
)

func main() {
	pkg.InitDB()

	router := gin.Default()

	router.Use(pkg.RateLimiter())

	router.POST("/login", pkg.Login)

	protected := router.Group("/")
	protected.Use(pkg.AuthMiddleware())
	{
		protected.PUT("/orders/:id", pkg.UpdateOrder)
		protected.DELETE("/orders/:id", pkg.DeleteOrder)
		protected.POST("/orders", pkg.CreateOrder)
	}

	router.GET("/orders", pkg.GetOrders)

	router.GET("/products", pkg.GetProducts)

	router.GET("/shipments", pkg.GetShipments)

	router.GET("/orders/:id", pkg.GetOrderByID)

	router.GET("/hi", pkg.SayHello)

	router.Run(":8080")
}
