package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtKey = []byte("my_secret_key")

func generateToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

var users = []Credentials{
	{Username: "user1", Password: "12345"},
	{Username: "user2", Password: "secret"},
	{Username: "user3", Password: "8068"},
}

func login(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	found := false
	for _, user := range users {
		if creds.Username == user.Username && creds.Password == user.Password {
			found = true
			break
		}
	}
	if !found {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	token, err := generateToken(creds.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

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
