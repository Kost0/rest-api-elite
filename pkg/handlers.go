package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetOrders godoc
// @Summary Get all orders in HTML table
// @Description Returns HTML page with table of all orders
// @ID get-orders-html
// @Produce html
// @Success 200 {string} string "HTML page with orders table"
// @Failure 500 {string} string "Internal Server Error"
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

func GetOrdersJSON(c *gin.Context) {
	var orders []Order
	db.Find(&orders)
	LogAction(c, "GET_ORDERS", 200)
	c.JSON(http.StatusOK, orders)
}

func GetProductsJSON(c *gin.Context) {
	var products []Product
	db.Find(&products)
	LogAction(c, "GET_PRODUCTS", 200)
	c.JSON(http.StatusOK, products)
}

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

func GetShipmentsJSON(c *gin.Context) {
	var shipments []Shipment
	db.Find(&shipments)
	LogAction(c, "GET_SHIPMENTS", 200)
	c.JSON(http.StatusOK, shipments)
}

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

func CreateOrder(c *gin.Context) {
	var newOrder Order
	if err := c.BindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		LogAction(c, "INVALID_REQUEST", 400)
		return
	}
	db.Create(&newOrder)
	LogAction(c, "CREATE_ORDER", 200)
	c.JSON(http.StatusCreated, newOrder)
}

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
