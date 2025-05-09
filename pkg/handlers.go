package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetOrders godoc
// @Summary Get all orders in HTML table
// @Description Returns HTML page with table of all orders
// @Tags HTML
// @ID get-orders-html
// @Produce html
// @Success 200 "HTML page with orders table"
// @Router /orders [get]
func GetOrders(c *gin.Context) {
	var orders []Order
	db.Find(&orders)
	LogAction(c, "GET_ORDERS_HTML", 200)
	tableView := TableView{
		Title:     "Заказы",
		Data:      orders,
		GetKeys:   GetKeys,
		GetValues: GetValues,
	}
	c.HTML(http.StatusOK, "index.html", tableView)
}

// GetOrdersJSON godoc
// @Summary Get all orders in JSON
// @Description Returns JSON with all orders
// @Tags Basic
// @ID get-orders-json
// @Produce json
// @Success 200 {array} Order "List of orders"
// @Router /JSON/orders [get]
func GetOrdersJSON(c *gin.Context) {
	var orders []Order
	db.Find(&orders)
	LogAction(c, "GET_ORDERS", 200)
	c.JSON(http.StatusOK, orders)
}

// GetProductsJSON godoc
// @Summary Get all products in JSON
// @Description Returns JSON with all products
// @Tags Basic
// @ID get-products-json
// @Produce json
// @Success 200 {array} Product "List of products"
// @Router /JSON/products [get]
func GetProductsJSON(c *gin.Context) {
	var products []Product
	db.Find(&products)
	LogAction(c, "GET_PRODUCTS", 200)
	c.JSON(http.StatusOK, products)
}

// GetProducts godoc
// @Summary Get all products in HTML table
// @Description Returns HTML page with table of all products
// @Tags HTML
// @ID get-products-html
// @Produce html
// @Success 200 "HTML page with products table"
// @Router /products [get]
func GetProducts(c *gin.Context) {
	var products []Product
	db.Find(&products)
	LogAction(c, "GET_PRODUCTS_HTML", 200)
	tableView := TableView{
		Title:     "Продукты",
		Data:      products,
		GetKeys:   GetKeys,
		GetValues: GetValues,
	}
	c.HTML(http.StatusOK, "index.html", tableView)
}

// GetShipmentsJSON godoc
// @Summary Get all shipments in JSON
// @Description Returns JSON with all shipments
// @Tags Basic
// @ID get-shipments-json
// @Produce json
// @Success 200 {array} Shipment "List of shipments"
// @Router /JSON/shipments [get]
func GetShipmentsJSON(c *gin.Context) {
	var shipments []Shipment
	db.Find(&shipments)
	LogAction(c, "GET_SHIPMENTS", 200)
	c.JSON(http.StatusOK, shipments)
}

// GetShipments godoc
// @Summary Get all Shipments in HTML table
// @Description Returns HTML page with table of all Shipments
// @Tags HTML
// @ID get-Shipments-html
// @Produce html
// @Success 200 "HTML page with Shipments table"
// @Router /shipments [get]
func GetShipments(c *gin.Context) {
	var shipments []Shipment
	db.Find(&shipments)
	LogAction(c, "GET_SHIPMENTS_HTML", 200)
	tableView := TableView{
		Title:     "Разгрузки",
		Data:      shipments,
		GetKeys:   GetKeys,
		GetValues: GetValues,
	}
	c.HTML(http.StatusOK, "index.html", tableView)
}

// GetOrderByID godoc
// @Summary Get order by its id
// @Description Return one order with certain id
// @Tags Basic
// @ID get-orders-by-id
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} Order
// @Failure 404 "Order not found"
// @Router /orders/:id [get]
func GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	var order Order
	if err := db.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
		LogAction(c, "ID_ABSENT", 404)
		return
	}
	LogAction(c, "GET_ORDER_BY_ID", 200)
	c.JSON(http.StatusOK, order)
}

// CreateOrder godoc
// @Summary Create new order
// @Description Make new order and add it to database
// @Tags Protected
// @ID create-order
// @Accept json
// @Param order body Order true "Order data in JSON"
// @Success 201 {object} Order
// @Failure 400 {string} Invalid request
// @Router /orders [post]
func CreateOrder(c *gin.Context) {
	var newOrder Order
	if err := c.BindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		LogAction(c, "INVALID_REQUEST", 400)
		return
	}
	db.Create(&newOrder)
	LogAction(c, "CREATE_ORDER", 201)
	c.JSON(http.StatusCreated, newOrder)
}

// UpdateOrder godoc
// @Summary Update order
// @Description Update order by its id
// @Tags Protected
// @ID update-order
// @Accept json
// @Param order body Order true "Order data in JSON"
// @Param id path string true "Order id"
// @Success 200 {object} Order
// @Failure 400 "Invalid request"
// @Failure 404	"Order not found"
// @Router /orders [put]
func UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	var updatedOrder Order
	if err := c.BindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		LogAction(c, "INVALID_REQUEST", 400)
		return
	}
	if err := db.Model(&Order{}).Where("id = ?", id).Updates(updatedOrder).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
		LogAction(c, "ID_ABSENT", 404)
		return
	}
	LogAction(c, "UPDATE_ORDER", 200)
	c.JSON(http.StatusOK, updatedOrder)
}

// DeleteOrder godoc
// @Summary Delete order
// @Description Delete order by its id
// @Tags Protected
// @ID delete-order
// @Param id path string true "Order id"
// @Success 200 "no content"
// @Failure 404 "Order not found"
// @Router /orders/:id [delete]
func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&Order{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
		LogAction(c, "ID_ABSENT", 404)
		return
	}
	LogAction(c, "DELETE_ORDER", 200)
	c.JSON(http.StatusOK, gin.H{"message": "order deleted"})
}

func SayHello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello")
}
