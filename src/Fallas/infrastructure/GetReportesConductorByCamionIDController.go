// GetReportesConductorByCamionIDController.go
package infrastructure

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetReportesConductorByCamionIDController struct {
	getByCamionIDUseCase *application.GetReportesConductorByCamionIDUseCase
}

func NewGetReportesConductorByCamionIDController(getByCamionIDUseCase *application.GetReportesConductorByCamionIDUseCase) *GetReportesConductorByCamionIDController {
	return &GetReportesConductorByCamionIDController{
		getByCamionIDUseCase: getByCamionIDUseCase,
	}
}

// @Summary      Reportes conductor por camión
// @Tags         ReporteConductor
// @Produce      json
// @Success      200 {object} entities.ReporteConductorResponse
// @Failure      400 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/reportes-conductor/camion/{camion_id} [get]
func (ctrl *GetReportesConductorByCamionIDController) Run(c *gin.Context) {
	camionIDParam := c.Param("camion_id")
	camionID, err := strconv.Atoi(camionIDParam)
	if err != nil {
		core.RespondBadRequest(c, "ID de camión inválido", map[string]string{"error": err.Error()})
		return
	}

	reportes, err := ctrl.getByCamionIDUseCase.Run(int32(camionID))
	if err != nil {
		core.RespondInternalServerError(c, "No se pudieron obtener los reportes del camión", err)
		return
	}

	core.RespondOK(c, reportes)
}
