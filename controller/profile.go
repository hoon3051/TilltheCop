package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoon3051/TilltheCop/form"
	"github.com/hoon3051/TilltheCop/service"
	"github.com/hoon3051/TilltheCop/util"
)

type ProfileController struct{}

var profileService service.ProfileService


func (ctr ProfileController) GetProfile(c *gin.Context) {
	// Get id from token
	userID, err := util.GetUserIDFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get profile (service)
	profile, err := profileService.GetProfile(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": profile})

}

func (ctr ProfileController) UpdateProfile(c *gin.Context) {
	// Get id from token
	userID, err := util.GetUserIDFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get the JSON body and decode into struct
	var profileForm form.ProfileForm
	if err := c.ShouldBindJSON(&profileForm); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validate the form
	if validationError := profileForm.Validate(); validationError != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": validationError})
		return
	}

	// Update profile (service)
	profile, err := profileService.UpdateProfile(userID, profileForm) 
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"Profile updated successfully", "profile": profile})

}
