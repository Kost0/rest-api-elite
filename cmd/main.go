package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"rest_api1/pkg"
)

// @title REST_API
// @version 1.0
// @description This is a sample server for delivery services.
// @termsOfService http://swagger.io/terms/
//
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
//
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
//
// @host localhost:8080
// @BasePath /rest_api1
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
