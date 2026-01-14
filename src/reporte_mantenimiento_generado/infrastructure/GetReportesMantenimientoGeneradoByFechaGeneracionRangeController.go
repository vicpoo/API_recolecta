// GetReportesMantenimientoGeneradoByFechaGeneracionRangeController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/reporte_mantenimiento_generado/application"
)

type GetReportesMantenimientoGeneradoByFechaGeneracionRangeController struct {
	getByFechaGeneracionRangeUseCase *application.GetReportesMantenimientoGeneradoByFechaGeneracionRangeUseCase
}

func NewGetReportesMantenimientoGeneradoByFechaGeneracionRangeController(getByFechaGeneracionRangeUseCase *application.GetReportesMantenimientoGeneradoByFechaGeneracionRangeUseCase) *GetReportesMantenimientoGeneradoByFechaGeneracionRangeController {
	return &GetReportesMantenimientoGeneradoByFechaGeneracionRangeController{
		getByFechaGeneracionRangeUseCase: getByFechaGeneracionRangeUseCase,
	}
}

func (ctrl *GetReportesMantenimientoGeneradoByFechaGeneracionRangeController) Run(c *gin.Context) {
	fechaInicio := c.Query("fecha_inicio")
	fechaFin := c.Query("fecha_fin")

	if fechaInicio == "" || fechaFin == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Se requieren ambos parámetros: fecha_inicio y fecha_fin",
		})
		return
	}

	reportes, err := ctrl.getByFechaGeneracionRangeUseCase.Run(fechaInicio, fechaFin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los reportes del rango de fechas de generación",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reportes)
}