// GetReportesFallaCriticaByFechaRangeController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/reporte_falla_critica/application"
)

type GetReportesFallaCriticaByFechaRangeController struct {
	getByFechaRangeUseCase *application.GetReportesFallaCriticaByFechaRangeUseCase
}

func NewGetReportesFallaCriticaByFechaRangeController(getByFechaRangeUseCase *application.GetReportesFallaCriticaByFechaRangeUseCase) *GetReportesFallaCriticaByFechaRangeController {
	return &GetReportesFallaCriticaByFechaRangeController{
		getByFechaRangeUseCase: getByFechaRangeUseCase,
	}
}

func (ctrl *GetReportesFallaCriticaByFechaRangeController) Run(c *gin.Context) {
	fechaInicio := c.Query("fecha_inicio")
	fechaFin := c.Query("fecha_fin")

	if fechaInicio == "" || fechaFin == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Se requieren los parámetros fecha_inicio y fecha_fin",
		})
		return
	}

	reportes, err := ctrl.getByFechaRangeUseCase.Run(fechaInicio, fechaFin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los reportes de falla crítica para el rango de fechas",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reportes)
}