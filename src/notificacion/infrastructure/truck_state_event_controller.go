package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

type TruckStateEventController struct {
	uc *application.ProcessTruckStateEventUseCase
}

func NewTruckStateEventController(uc *application.ProcessTruckStateEventUseCase) *TruckStateEventController {
	return &TruckStateEventController{uc: uc}
}

func (ctrl *TruckStateEventController) Process(c *gin.Context) {
	var event domain.TruckStateEvent
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload invalido", "details": err.Error()})
		return
	}

	output, err := ctrl.uc.Execute(c.Request.Context(), &event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}
