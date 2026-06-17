// GetSeguimientosFallaCriticaByFechaRangeController.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetSeguimientosFallaCriticaByFechaRangeController struct {
	getByFechaRangeUseCase *application.GetSeguimientosFallaCriticaByFechaRangeUseCase
}

func NewGetSeguimientosFallaCriticaByFechaRangeController(getByFechaRangeUseCase *application.GetSeguimientosFallaCriticaByFechaRangeUseCase) *GetSeguimientosFallaCriticaByFechaRangeController {
	return &GetSeguimientosFallaCriticaByFechaRangeController{
		getByFechaRangeUseCase: getByFechaRangeUseCase,
	}
}

// @Summary      Seguimientos falla crítica por fecha
// @Tags         SeguimientoFallaCritica
// @Produce      json
// @Success      200 {object} entities.SeguimientoFallaCriticaResponse
// @Failure      400 {object} core.ErrorResponse
// @Router       /api/seguimientos-falla-critica/por-fecha [get]
func (ctrl *GetSeguimientosFallaCriticaByFechaRangeController) Run(c *gin.Context) {
	fechaInicio := c.Query("fecha_inicio")
	fechaFin := c.Query("fecha_fin")

	if fechaInicio == "" || fechaFin == "" {
		core.RespondBadRequest(c, "Se requieren los parámetros fecha_inicio y fecha_fin", nil)
		return
	}

	seguimientos, err := ctrl.getByFechaRangeUseCase.Run(fechaInicio, fechaFin)
	if err != nil {
		core.RespondInternalServerError(c, "No se pudieron obtener los seguimientos para el rango de fechas", err)
		return
	}

	core.RespondOK(c, seguimientos)
}
