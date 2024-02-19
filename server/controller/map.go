package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoon3051/TilltheCop/server/form"
	"github.com/hoon3051/TilltheCop/server/service"
)

type MapController struct{}

var mapService service.MapService

func (ctr MapController) GetMap(c *gin.Context) {
	// Get location from body
	var locationForm form.LocationForm
	if err := c.ShouldBindJSON(&locationForm); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate the form
	if validationError := locationForm.Validate(); validationError != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": validationError})
		return
	}

	// Get map (service)
	mapURL, err := mapService.GetMap(locationForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//c.Redirect(http.StatusTemporaryRedirect, mapURL)
	c.JSON(http.StatusOK, gin.H{"map_url": mapURL})
}
