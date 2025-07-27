package middleware

import (
	"donasi-yuk/helper"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: missing bearer token"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		userID, err := helper.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: invalid token"})
			return
		}

		// Simpan user ID ke context
		c.Set("userID", userID)
		c.Next()

		fmt.Println("Token received:", tokenString)
		fmt.Println("UserID from token:", userID)
	}

}
