// GetReportesFallaCriticaByCamionIDController.go
package infrastructure

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetReportesFallaCriticaByCamionIDController struct {
	getByCamionIDUseCase *application.GetReportesFallaCriticaByCamionIDUseCase
}

func NewGetReportesFallaCriticaByCamionIDController(getByCamionIDUseCase *application.GetReportesFallaCriticaByCamionIDUseCase) *GetReportesFallaCriticaByCamionIDController {
	return &GetReportesFallaCriticaByCamionIDController{
		getByCamionIDUseCase: getByCamionIDUseCase,
	}
}

// @Summary      Reportes falla crítica por camión
// @Tags         ReporteFallaCritica
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/reportes-falla-critica/camion/{camionId} [get]
func (ctrl *GetReportesFallaCriticaByCamionIDController) Run(c *gin.Context) {
	camionIDParam := c.Param("camionId")
	camionID, err := strconv.Atoi(camionIDParam)
	if err != nil {
		core.RespondInvalidInput(c, "ID de camión inválido")
		return
	}

	reportes, err := ctrl.getByCamionIDUseCase.Run(int32(camionID))
	if err != nil {
		core.RespondInternalServerError(c, "No se pudieron obtener los reportes de falla crítica para el camión", err)
		return
	}

	core.RespondOK(c, reportes)
}
