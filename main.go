package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=221706 dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Миграция схемы
	db.AutoMigrate(&Order{})
}

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

type Order struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Customer     string `json:"customer"`
	Amount       int    `json:"amount"`
	Address      string `json:"address"`
	Code         int    `json:"code"`
	Phone        string `json:"phone"`
	Product_name string `json:"product_name"`
}

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

func getOrders(c *gin.Context) {
	var orders []Order
	db.Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func getOrderByID(c *gin.Context) {
	id := c.Param("id")
	var book Order
	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func createOrder(c *gin.Context) {
	var newBook Order
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	db.Create(&newBook)
	c.JSON(http.StatusCreated, newBook)
}

func updateOrder(c *gin.Context) {
	id := c.Param("id")
	var updatedBook Order
	if err := c.BindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	if err := db.Model(&Order{}).Where("id = ?", id).Updates(updatedBook).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
		return
	}
	c.JSON(http.StatusOK, updatedBook)
}

func deleteOrder(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&Order{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "order deleted"})
}
