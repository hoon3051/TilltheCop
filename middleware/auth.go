package middleware

import (
	"net/http"

	"github.com/hoon3051/TilltheCop/service"

	"github.com/gin-gonic/gin"
)

func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from header
		tokenString := c.GetHeader("Authorization")

		// Check token
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized123",
			})
			return
		}

		// Get userID from token
		userID, err := service.AuthService{}.ExtractTokenID(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Set userID to context
		c.Set("user_id", userID)

		// Next
		c.Next()
	}
}
