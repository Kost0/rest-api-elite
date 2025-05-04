package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"rest_api1/pkg"
)

func main() {
	pkg.InitDB()

	router := gin.Default()

	router.Use(pkg.RateLimiter())

	router.SetFuncMap(template.FuncMap{
		"GetKeys":   pkg.GetKeys,
		"GetValues": pkg.GetValues,
	})

	router.LoadHTMLGlob("pkg/templates/*")

	router.POST("/login", pkg.Login)

	protected := router.Group("/")
	protected.Use(pkg.AuthMiddleware())
	{
		protected.PUT("/orders/:id", pkg.UpdateOrder)
		protected.DELETE("/orders/:id", pkg.DeleteOrder)
		protected.POST("/orders", pkg.CreateOrder)
	}

	jsons := router.Group("/JSON/")

	jsons.GET("/orders", pkg.GetOrdersJSON)

	jsons.GET("/products", pkg.GetProductsJSON)

	jsons.GET("/shipments", pkg.GetShipmentsJSON)

	router.GET("/orders", pkg.GetOrders)

	router.GET("/products", pkg.GetProducts)

	router.GET("/shipments", pkg.GetShipments)

	router.GET("/orders/:id", pkg.GetOrderByID)

	router.GET("/hi", pkg.SayHello)

	router.Run(":8080")
}
