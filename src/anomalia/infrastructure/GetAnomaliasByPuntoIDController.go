// GetAnomaliasByPuntoIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/anomalia/application"
)

type GetAnomaliasByPuntoIDController struct {
	getByPuntoIDUseCase *application.GetAnomaliasByPuntoIDUseCase
}

func NewGetAnomaliasByPuntoIDController(getByPuntoIDUseCase *application.GetAnomaliasByPuntoIDUseCase) *GetAnomaliasByPuntoIDController {
	return &GetAnomaliasByPuntoIDController{
		getByPuntoIDUseCase: getByPuntoIDUseCase,
	}
}

func (ctrl *GetAnomaliasByPuntoIDController) Run(c *gin.Context) {
	puntoIDParam := c.Param("puntoId")
	puntoID, err := strconv.Atoi(puntoIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de punto inválido",
			"error":   err.Error(),
		})
		return
	}

	anomalias, err := ctrl.getByPuntoIDUseCase.Run(int32(puntoID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las anomalías para el punto",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, anomalias)
}