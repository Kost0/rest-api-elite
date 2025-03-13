package main

import "github.com/dgrijalva/jwt-go"

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
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
