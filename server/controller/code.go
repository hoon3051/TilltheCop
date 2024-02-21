package controller

import (
	"net/http"

	"github.com/hoon3051/TilltheCop/server/service"
	"github.com/hoon3051/TilltheCop/server/util"

	"github.com/gin-gonic/gin"
)

type CodeController struct{}

var codeService service.CodeService

func (ctr CodeController) GenerateQRCode(c *gin.Context) {
	userID, err := util.GetUserIDFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	reportID := c.Param("reportID")
	if reportID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "reportID is required"})
		return
	}

	qrCode, err := codeService.GenerateQRCode(userID, reportID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "image/png", qrCode)

}
