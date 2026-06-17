// GetAlertasByFechaRangeController.go
package infrastructure

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetAlertasByFechaRangeController struct {
	getByFechaRangeUseCase *application.GetAlertasByFechaRangeUseCase
}

func NewGetAlertasByFechaRangeController(getByFechaRangeUseCase *application.GetAlertasByFechaRangeUseCase) *GetAlertasByFechaRangeController {
	return &GetAlertasByFechaRangeController{
		getByFechaRangeUseCase: getByFechaRangeUseCase,
	}
}

// @Summary      Alertas por rango de fecha
// @Tags         AlertaMantenimiento
// @Produce      json
// @Success      200 {object} entities.AlertaMantenimientoResponse
// @Failure      400 {object} core.ErrorResponse
// @Router       /api/alertas-mantenimiento/fecha [get]
func (ctrl *GetAlertasByFechaRangeController) Run(c *gin.Context) {
	var request struct {
		FechaInicio string `form:"fecha_inicio" binding:"required"`
		FechaFin    string `form:"fecha_fin" binding:"required"`
	}

	if err := c.ShouldBindQuery(&request); err != nil {
		core.RespondValidationError(c, "Parámetros de fecha requeridos", map[string]string{
			"params": "fecha_inicio, fecha_fin",
			"error":  err.Error(),
		})
		return
	}

	// Parsear fechas
	layout := "2006-01-02"
	fechaInicio, err := time.Parse(layout, request.FechaInicio)
	if err != nil {
		core.RespondBadRequest(c, "Formato de fecha_inicio inválido. Use YYYY-MM-DD", map[string]string{
			"field": "fecha_inicio",
			"error": err.Error(),
		})
		return
	}

	fechaFin, err := time.Parse(layout, request.FechaFin)
	if err != nil {
		core.RespondBadRequest(c, "Formato de fecha_fin inválido. Use YYYY-MM-DD", map[string]string{
			"field": "fecha_fin",
			"error": err.Error(),
		})
		return
	}

	// Ajustar fecha fin para incluir todo el día
	fechaFin = fechaFin.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	alertas, err := ctrl.getByFechaRangeUseCase.Run(fechaInicio, fechaFin)
	if err != nil {
		core.RespondInternalServerError(c, "Error al obtener alertas por rango de fechas", err)
		return
	}

	core.RespondOK(c, map[string]interface{}{
		"data": alertas,
	})
}
