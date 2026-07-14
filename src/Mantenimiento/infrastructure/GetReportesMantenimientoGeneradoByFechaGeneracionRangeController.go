// GetReportesMantenimientoGeneradoByFechaGeneracionRangeController.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetReportesMantenimientoGeneradoByFechaGeneracionRangeController struct {
	getByFechaGeneracionRangeUseCase *application.GetReportesMantenimientoGeneradoByFechaGeneracionRangeUseCase
}

func NewGetReportesMantenimientoGeneradoByFechaGeneracionRangeController(getByFechaGeneracionRangeUseCase *application.GetReportesMantenimientoGeneradoByFechaGeneracionRangeUseCase) *GetReportesMantenimientoGeneradoByFechaGeneracionRangeController {
	return &GetReportesMantenimientoGeneradoByFechaGeneracionRangeController{
		getByFechaGeneracionRangeUseCase: getByFechaGeneracionRangeUseCase,
	}
}

// @Summary      Reportes por fecha de generación
// @Tags         ReporteMantenimientoGenerado
// @Produce      json
// @Success      200 {object} entities.RegistroMantenimientoResponse
// @Failure      400 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/reportes-mantenimiento-generado/fecha-generacion [get]
func (ctrl *GetReportesMantenimientoGeneradoByFechaGeneracionRangeController) Run(c *gin.Context) {
	fechaInicio := c.Query("fecha_inicio")
	fechaFin := c.Query("fecha_fin")

	if fechaInicio == "" || fechaFin == "" {
		core.RespondBadRequest(c, "Se requieren ambos parámetros: fecha_inicio y fecha_fin", map[string]string{
			"fecha_inicio": "requerido",
			"fecha_fin":    "requerido",
		})
		return
	}

	reportes, err := ctrl.getByFechaGeneracionRangeUseCase.Run(fechaInicio, fechaFin)
	if err != nil {
		core.RespondInternalServerError(c, "No se pudieron obtener los reportes del rango de fechas de generación", err)
		return
	}

	core.RespondOK(c, map[string]interface{}{
		"data": reportes,
	})
}
