// GetReportesConductorByRutaIDController.go
package infrastructure

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetReportesConductorByRutaIDController struct {
	getByRutaIDUseCase *application.GetReportesConductorByRutaIDUseCase
}

func NewGetReportesConductorByRutaIDController(getByRutaIDUseCase *application.GetReportesConductorByRutaIDUseCase) *GetReportesConductorByRutaIDController {
	return &GetReportesConductorByRutaIDController{
		getByRutaIDUseCase: getByRutaIDUseCase,
	}
}

// @Summary      Reportes conductor por ruta
// @Tags         ReporteConductor
// @Produce      json
// @Success      200 {object} entities.ReporteConductorResponse
// @Failure      400 {object} core.ErrorResponse
// @Router       /api/reportes-conductor/ruta/{ruta_id} [get]
func (ctrl *GetReportesConductorByRutaIDController) Run(c *gin.Context) {
	rutaIDParam := c.Param("ruta_id")
	rutaID, err := strconv.Atoi(rutaIDParam)
	if err != nil {
		core.RespondBadRequest(c, "ID de ruta inválido", map[string]string{"error": err.Error()})
		return
	}

	reportes, err := ctrl.getByRutaIDUseCase.Run(int32(rutaID))
	if err != nil {
		core.RespondInternalServerError(c, "No se pudieron obtener los reportes de la ruta", err)
		return
	}

	core.RespondOK(c, reportes)
}
