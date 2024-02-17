package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/hoon3051/TilltheCop/service"
	"github.com/hoon3051/TilltheCop/form"
	"github.com/hoon3051/TilltheCop/initializer"

)

type OauthController struct{}

var oauthService service.OauthService

func (ctr OauthController) GoogleLogin(c *gin.Context) {
	// Generate state string
	state, err := oauthService.GenerateStateString()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate state string"})
		return
	}

	//Save state string in session
	session := sessions.Default(c)
	session.Set("oauthState", state)
	err = session.Save()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid state"})
		return
	}

	// Handle the exchange code to initiate a transport.
	code := c.Query("code")

	// Get token (service)
	oauthToken, err := oauthService.GetOauthToken(code)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get userInfo (service)
	userInfo, err := oauthService.GetOauthUser(oauthToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check whether the user exists (service)
	userExists := oauthService.FindUserExists(userInfo)

	if !userExists { // if user does not exist (Register service)
		// Save oauthToken and userInfo in session
		tokenJson, err := json.Marshal(oauthToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize oauthToken"})
			return
		}
		session.Set("oauthToken", string(tokenJson))

		userInfoJson, err := json.Marshal(userInfo)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize userInfo"})
			return
		}
		session.Set("userInfo", string(userInfoJson))

		err = session.Save()
		if err != nil {
			log.Printf("Session save error: %v", err) // 로깅 추가
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session1"})
			return
		}

		// Redirect to create user page
		c.Redirect(http.StatusTemporaryRedirect, "/oauth/register")

	}
	// if user exists (Login)

	// Save user in DB (service)
	user, token, err := oauthService.SaveOauthUser(oauthToken, userInfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Clear state from session
	session.Delete("oauthState")
	err = session.Save()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session2"})
		return
	}

	// Return the user and token
	c.JSON(http.StatusOK, gin.H{"message": "Logged in Success", "user": user, "oauthToken": oauthToken, "token": token})

}

func (ctr OauthController) Register(c *gin.Context) {
	// Get oauthToken and userInfo from session
	session := sessions.Default(c)
	oauthTokenString := session.Get("oauthToken")
	userInfoString := session.Get("userInfo")

	// Check session
	if oauthTokenString == nil || userInfoString == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to get session"})
		return
	}

	// Parse oauthToken and userInfo
	var oauthToken form.OauthToken
	err := json.Unmarshal([]byte(oauthTokenString.(string)), &oauthToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse oauthToken"})
		return
	}

	var userInfo form.OauthUser
	err = json.Unmarshal([]byte(userInfoString.(string)), &userInfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse userInfo"})
		return
	}

	// Start transaction
	tx := initializer.DB.Begin()

	// Register user in DB (service)
	var userService service.UserService
	userID, err := userService.Register(tx, userInfo)
	if err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Register oauth in DB (service)
	err = oauthService.Register(tx, oauthToken, userID)
	if err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get the JSON body and decode into struct
	var profileForm form.ProfileForm
	if err := c.ShouldBindJSON(&profileForm); err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validate the form
	if validationError := profileForm.Validate(); validationError != "" {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": validationError})
		return
	}

	// Create profile in DB (service) 
	err = profileService.CreateProfile(tx, profileForm, userID)
	if err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Commit transaction
	tx.Commit()

	// Clear session
	session.Delete("oauthToken")
	session.Delete("userInfo")
	err = session.Save()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	// Redirect to login page
	c.Redirect(http.StatusTemporaryRedirect, "/oauth/google/login")

}