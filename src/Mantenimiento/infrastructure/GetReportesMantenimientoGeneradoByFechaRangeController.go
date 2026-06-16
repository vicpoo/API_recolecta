// GetReportesMantenimientoGeneradoByFechaRangeController.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/core"
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
		core.RespondBadRequest(c, "Se requieren ambos parámetros: fecha_inicio y fecha_fin", map[string]string{
			"fecha_inicio": "requerido",
			"fecha_fin":    "requerido",
		})
		return
	}

	reportes, err := ctrl.getByFechaRangeUseCase.Run(fechaInicio, fechaFin)
	if err != nil {
		core.RespondInternalServerError(c, "No se pudieron obtener los reportes del rango de fechas", err)
		return
	}

	core.RespondOK(c, map[string]interface{}{
		"data": reportes,
	})
}
