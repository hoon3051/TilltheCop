package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/hoon3051/TilltheCop/server/form"
	"github.com/hoon3051/TilltheCop/server/service"
)

type AuthController struct {}

var tokenService = service.AuthService{}

func(ctr AuthController) RefreshToken(c *gin.Context) {
	// Refresh Token을 요청 본문에서 가져옵니다. (JSON)
	var refreshtoken form.RefreshTokenForm
	if err := c.ShouldBindJSON(&refreshtoken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid refreshToken json provided"})
		return
	}

	// Validate the token
	if validationError := refreshtoken.Validate(); validationError != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationError})
		return
	}

	// Parse the token
	claims := &jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(refreshtoken.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})

	// Handle errors from parsing
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	// Check if token is expired
	if time.Unix(claims.ExpiresAt.Time.Unix(), 0).Before(time.Now()) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token expired"})
		return
	}

	// Create new Access Token (service)
	userID, err := tokenService.ExtractTokenID(refreshtoken.RefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting user ID from token"})
		return
	}

	tokenDetails, err := tokenService.CreateToken(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating new access token"})
		return
	}

	// Save the details of the new token in Redis (service)
	err = tokenService.SaveToken(userID, tokenDetails)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving the new token details"})
		return
	}

	// Return the new tokens
	c.JSON(http.StatusOK, gin.H{"access_token": tokenDetails.AccessToken})

}