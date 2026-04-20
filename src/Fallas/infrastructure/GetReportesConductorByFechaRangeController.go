// GetReportesConductorByFechaRangeController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
)

type GetReportesConductorByFechaRangeController struct {
	getByFechaRangeUseCase *application.GetReportesConductorByFechaRangeUseCase
}

func NewGetReportesConductorByFechaRangeController(getByFechaRangeUseCase *application.GetReportesConductorByFechaRangeUseCase) *GetReportesConductorByFechaRangeController {
	return &GetReportesConductorByFechaRangeController{
		getByFechaRangeUseCase: getByFechaRangeUseCase,
	}
}

// @Summary      Reportes conductor por fecha
// @Tags         ReporteConductor
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/reportes-conductor/fecha [get]
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