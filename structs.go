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
	ID            uint   `gorm:"primaryKey" json:"id"`
	Amount        int    `json:"amount"`
	Address       string `json:"address"`
	Code          int    `json:"code"`
	Phone         string `json:"phone"`
	Product_name  string `json:"product_name"`
	Customer      int    `json:"customer"`
	Delivery_team int    `json:"delivery_team"`
}

type Product struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Product_name string `json:"product_name"`
	Price        int    `json:"price"`
	Supplier_id  int    `json:"supplier_id"`
}

type Shipment struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Order_id      int    `json:"order_id"`
	Shipment_date string `json:"shipment_date"`
	amount        int    `json:"amount"`
}
