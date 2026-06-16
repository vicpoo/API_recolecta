// GetReportesFallaCriticaByFechaRangeController.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetReportesFallaCriticaByFechaRangeController struct {
	getByFechaRangeUseCase *application.GetReportesFallaCriticaByFechaRangeUseCase
}

func NewGetReportesFallaCriticaByFechaRangeController(getByFechaRangeUseCase *application.GetReportesFallaCriticaByFechaRangeUseCase) *GetReportesFallaCriticaByFechaRangeController {
	return &GetReportesFallaCriticaByFechaRangeController{
		getByFechaRangeUseCase: getByFechaRangeUseCase,
	}
}

// @Summary      Reportes falla crítica por fecha
// @Tags         ReporteFallaCritica
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/reportes-falla-critica/por-fecha [get]
func (ctrl *GetReportesFallaCriticaByFechaRangeController) Run(c *gin.Context) {
	fechaInicio := c.Query("fecha_inicio")
	fechaFin := c.Query("fecha_fin")

	if fechaInicio == "" || fechaFin == "" {
		core.RespondBadRequest(c, "Se requieren los parámetros fecha_inicio y fecha_fin", nil)
		return
	}

	reportes, err := ctrl.getByFechaRangeUseCase.Run(fechaInicio, fechaFin)
	if err != nil {
		core.RespondInternalServerError(c, "No se pudieron obtener los reportes de falla crítica para el rango de fechas", err)
		return
	}

	core.RespondOK(c, reportes)
}
