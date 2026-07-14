// GetReportesConductorByFechaRangeController.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/core"
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
// @Success      200 {object} entities.ReporteConductorResponse
// @Failure      400 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/reportes-conductor/fecha [get]
func (ctrl *GetReportesConductorByFechaRangeController) Run(c *gin.Context) {
	fechaInicio := c.Query("fecha_inicio")
	fechaFin := c.Query("fecha_fin")

	if fechaInicio == "" || fechaFin == "" {
		core.RespondBadRequest(c, "Se requieren ambos parámetros: fecha_inicio y fecha_fin", nil)
		return
	}

	reportes, err := ctrl.getByFechaRangeUseCase.Run(fechaInicio, fechaFin)
	if err != nil {
		core.RespondInternalServerError(c, "No se pudieron obtener los reportes del rango de fechas", err)
		return
	}

	core.RespondOK(c, reportes)
}
