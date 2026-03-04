// GetReportesMantenimientoGeneradoByFechaRangeController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
)

type GetReportesMantenimientoGeneradoByFechaRangeController struct {
	getByFechaRangeUseCase *application.GetReportesMantenimientoGeneradoByFechaRangeUseCase
}

func NewGetReportesMantenimientoGeneradoByFechaRangeController(getByFechaRangeUseCase *application.GetReportesMantenimientoGeneradoByFechaRangeUseCase) *GetReportesMantenimientoGeneradoByFechaRangeController {
	return &GetReportesMantenimientoGeneradoByFechaRangeController{
		getByFechaRangeUseCase: getByFechaRangeUseCase,
	}
}

// @Summary      Reportes por rango de fecha
// @Tags         ReporteMantenimientoGenerado
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/reportes-mantenimiento-generado/fecha [get]
func (ctrl *GetReportesMantenimientoGeneradoByFechaRangeController) Run(c *gin.Context) {
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