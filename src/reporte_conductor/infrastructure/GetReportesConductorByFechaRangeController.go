// GetReportesConductorByFechaRangeController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/reporte_conductor/application"
)

type GetReportesConductorByFechaRangeController struct {
	getByFechaRangeUseCase *application.GetReportesConductorByFechaRangeUseCase
}

func NewGetReportesConductorByFechaRangeController(getByFechaRangeUseCase *application.GetReportesConductorByFechaRangeUseCase) *GetReportesConductorByFechaRangeController {
	return &GetReportesConductorByFechaRangeController{
		getByFechaRangeUseCase: getByFechaRangeUseCase,
	}
}

func (ctrl *GetReportesConductorByFechaRangeController) Run(c *gin.Context) {
	fechaInicio := c.Query("fecha_inicio")
	fechaFin := c.Query("fecha_fin")

	if fechaInicio == "" || fechaFin == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Se requieren ambos parámetros: fecha_inicio y fecha_fin",
		})
		return
	}

	reportes, err := ctrl.getByFechaRangeUseCase.Run(fechaInicio, fechaFin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los reportes del rango de fechas",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reportes)
}