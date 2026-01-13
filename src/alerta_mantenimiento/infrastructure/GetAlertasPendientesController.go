//GetAlertasPendientesController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/application"
)

type GetAlertasPendientesController struct {
	getPendientesUseCase *application.GetAlertasPendientesUseCase
}

func NewGetAlertasPendientesController(getPendientesUseCase *application.GetAlertasPendientesUseCase) *GetAlertasPendientesController {
	return &GetAlertasPendientesController{
		getPendientesUseCase: getPendientesUseCase,
	}
}

func (ctrl *GetAlertasPendientesController) Run(c *gin.Context) {
	alertas, err := ctrl.getPendientesUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las alertas pendientes",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, alertas)
}