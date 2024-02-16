package controller

import (
	"net/http"
	"encoding/json"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/hoon3051/TilltheCop/service"

)

type OauthController struct{}

var oauthService service.OauthService

func (ctr OauthController) GoogleLogin(c *gin.Context) {
	// Generate state string
	state, err := oauthService.GenerateStateString()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate state string"})
		return
	}

	//Save state string in session
	session := sessions.Default(c)
	session.Set("oauthState", state)
	err = session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	// Redirect user to Google's consent page
	url := oauthService.GetGoogleOauthURL(state)
	c.Redirect(http.StatusTemporaryRedirect, url)

}

func (ctr OauthController) GoogleCallback(c *gin.Context) {
	// Get state from session
	receivedState := c.Query("state")
	session := sessions.Default(c)
	savedState := session.Get("oauthState")

	// Check state
	if savedState == nil || savedState != receivedState {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state"})
		return
	}

	// Handle the exchange code to initiate a transport.
	code := c.Query("code")

	// Get token (service)
	oauthToken, err := oauthService.GetOauthToken(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get userInfo (service)
	userInfo, err := oauthService.GetOauthUser(oauthToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check whether the user exists (service)
	userExists := oauthService.FindUserExists(userInfo)

	if !userExists { // if user does not exist (Register service)
		// Save oauthToken and userInfo in session
		tokenJson, err := json.Marshal(oauthToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize oauthToken"})
			return
		}
		session.Set("oauthToken", string(tokenJson))

		userInfoJson, err := json.Marshal(userInfo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize userInfo"})
			return
		}
		session.Set("userInfo", string(userInfoJson))

		err = session.Save()
		if err != nil {
			log.Printf("Session save error: %v", err) // 로깅 추가
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session1"})
			return
		}

		// Register user in DB (service)
		var userService service.UserService
		userID, err := userService.Register(userInfo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Register oauth in DB (service)
		err = oauthService.Register(oauthToken, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	// if user exists (Login)

	// Save user in DB (service)
	user, token, err := oauthService.SaveOauthUser(oauthToken, userInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Clear state from session
	session.Delete("oauthState")
	err = session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session2"})
		return
	}

	// Return the user and token
	c.JSON(http.StatusOK, gin.H{"message":"Logged in Success","user": user, "oauthToken": oauthToken, "token": token})

}