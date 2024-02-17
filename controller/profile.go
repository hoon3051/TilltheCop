package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoon3051/TilltheCop/form"
	"github.com/hoon3051/TilltheCop/service"
)

type ProfileController struct{}

var profileService service.ProfileService


func (ctr ProfileController) GetProfile(c *gin.Context) {
	// Get id from session
	id, _ := c.Get("id")

	// Get profile (service)
	profile, err := profileService.GetProfile(id.(uint))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (ctr ProfileController) UpdateProfile(c *gin.Context) {
	// Get id from session
	id, _ := c.Get("id")

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
	profile, err := profileService.UpdateProfile(id.(uint), profileForm) 
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"Profile updated successfully", "profile": profile})

}
