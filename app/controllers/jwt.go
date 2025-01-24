package controllers

import (
	authservice "MATHB/app/services/authService"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RefreshJWT(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	// 解析Token
	claims, err := authservice.ParseJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// 检查Token是否接近过期（例如10分钟内）
	if time.Until(claims.ExpiresAt.Time) > time.Minute*10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is not close to expiration"})
		return
	}

	// 生成新的Token
	newToken, err := authservice.CreateJWT(claims.UserID, time.Hour*2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": newToken})
}
