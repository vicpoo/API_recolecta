// GetReportesFallaCriticaByConductorIDController.go
package infrastructure

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetReportesFallaCriticaByConductorIDController struct {
	getByConductorIDUseCase *application.GetReportesFallaCriticaByConductorIDUseCase
}

func NewGetReportesFallaCriticaByConductorIDController(getByConductorIDUseCase *application.GetReportesFallaCriticaByConductorIDUseCase) *GetReportesFallaCriticaByConductorIDController {
	return &GetReportesFallaCriticaByConductorIDController{
		getByConductorIDUseCase: getByConductorIDUseCase,
	}
}

// @Summary      Reportes falla crítica por conductor
// @Tags         ReporteFallaCritica
// @Produce      json
// @Success      200 {object} entities.ReporteFallaCriticaResponse
// @Failure      400 {object} core.ErrorResponse
// @Router       /api/reportes-falla-critica/conductor/{conductorId} [get]
func (ctrl *GetReportesFallaCriticaByConductorIDController) Run(c *gin.Context) {
	conductorIDParam := c.Param("conductorId")
	conductorID, err := strconv.Atoi(conductorIDParam)
	if err != nil {
		core.RespondInvalidInput(c, "ID de conductor inválido")
		return
	}

	reportes, err := ctrl.getByConductorIDUseCase.Run(int32(conductorID))
	if err != nil {
		core.RespondInternalServerError(c, "No se pudieron obtener los reportes de falla crítica para el conductor", err)
		return
	}

	core.RespondOK(c, reportes)
}
