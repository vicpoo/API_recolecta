// GetAnomaliasByFechaRangeController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/anomalia/application"
)

type GetAnomaliasByFechaRangeController struct {
	getByFechaRangeUseCase *application.GetAnomaliasByFechaRangeUseCase
}

func NewGetAnomaliasByFechaRangeController(getByFechaRangeUseCase *application.GetAnomaliasByFechaRangeUseCase) *GetAnomaliasByFechaRangeController {
	return &GetAnomaliasByFechaRangeController{
		getByFechaRangeUseCase: getByFechaRangeUseCase,
	}
}

func (ctrl *GetAnomaliasByFechaRangeController) Run(c *gin.Context) {
	fechaInicio := c.Query("fecha_inicio")
	fechaFin := c.Query("fecha_fin")

	if fechaInicio == "" || fechaFin == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Se requieren los parámetros fecha_inicio y fecha_fin",
		})
		return
	}

	anomalias, err := ctrl.getByFechaRangeUseCase.Run(fechaInicio, fechaFin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las anomalías para el rango de fechas",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, anomalias)
}