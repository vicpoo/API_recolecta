//GetAlertasByCamionIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/application"
)

type GetAlertasByCamionIDController struct {
	getByCamionUseCase *application.GetAlertasByCamionIDUseCase
}

func NewGetAlertasByCamionIDController(getByCamionUseCase *application.GetAlertasByCamionIDUseCase) *GetAlertasByCamionIDController {
	return &GetAlertasByCamionIDController{
		getByCamionUseCase: getByCamionUseCase,
	}
}

func (ctrl *GetAlertasByCamionIDController) Run(c *gin.Context) {
	idParam := c.Param("camion_id")
	camionID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de camión inválido",
			"error":   err.Error(),
		})
		return
	}

	alertas, err := ctrl.getByCamionUseCase.Run(int32(camionID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las alertas del camión",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, alertas)
}