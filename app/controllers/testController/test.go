package testcontroller

import (
	authservice "MATHB/app/services/authService"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TestHandler(c *gin.Context) {
	// 生成新的Token
	newToken, err := authservice.CreateJWT(123456789, time.Hour*2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": newToken})
}
