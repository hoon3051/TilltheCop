package controller

import (
	"net/http"

	"github.com/hoon3051/TilltheCop/server/form"
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

	reportID := c.Query("reportID")
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

func (ctr CodeController) ScanQRCode(c *gin.Context) {
	userID, err := util.GetUserIDFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var codeData form.CodeData
	if err := c.ShouldBindJSON(&codeData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record, err := codeService.CreateRecord(userID, codeData.ReportID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"record": record})
}
