package models

import "github.com/golang-jwt/jwt/v4"

type CustomClaims struct {
	UserID uint `json:"userID"`
	jwt.RegisteredClaims
}
