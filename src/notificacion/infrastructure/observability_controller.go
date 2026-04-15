package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
)

type ObservabilityController struct {
	uc *application.GetObservabilitySummaryUseCase
}

func NewObservabilityController(uc *application.GetObservabilitySummaryUseCase) *ObservabilityController {
	return &ObservabilityController{uc: uc}
}

func (ctrl *ObservabilityController) Summary(c *gin.Context) {
	truckIDRaw := c.Param("truck_id")
	truckIDParsed, err := strconv.ParseInt(truckIDRaw, 10, 32)
	if err != nil || truckIDParsed <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "truck_id invalido"})
		return
	}

	output, err := ctrl.uc.Execute(c.Request.Context(), int32(truckIDParsed))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}
