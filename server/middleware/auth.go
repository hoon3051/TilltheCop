package middleware

import (
	"net/http"
	"strings" // 이 부분을 추가

	"github.com/gin-gonic/gin"
	"github.com/hoon3051/TilltheCop/server/service"
)

func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from header
		authHeader := c.GetHeader("Authorization")

		// Check if Authorization header is present
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is missing.",
			})
			return
		}

		// Extract the token from the Authorization header
		tokenString := ""
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header format is not valid.",
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

		// Continue to the next middleware
		c.Next()
	}
}
