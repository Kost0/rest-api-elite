package pkg

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

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

// Login godoc
// @Summary Login in system
// @Description User must enter login and password to get token
// @Tags Login
// @ID login
// @Param data body Credentials true "Users credentials"
// @Success 200 {string} token
// @Failure 400 "Invalid request"
// @Failure 401 "Wrong data"
// @Failure 500 "Cant generate token"
// @Router /login [post]
func Login(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		LogAction(c, "INVALID_LOGIN_REQUEST", 400)
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
		LogAction(c, "WRONG_INFO", 401)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	token, err := generateToken(creds.Username)
	if err != nil {
		LogAction(c, "TOKEN_ERROR", 500)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create token"})
		return
	}
	LogAction(c, "SUCCESS_LOGIN", 200)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			LogAction(c, "WRONG_TOKEN", 401)
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
