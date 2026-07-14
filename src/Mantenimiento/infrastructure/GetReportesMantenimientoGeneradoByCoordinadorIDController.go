// GetReportesMantenimientoGeneradoByCoordinadorIDController.go
package infrastructure

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetReportesMantenimientoGeneradoByCoordinadorIDController struct {
	getByCoordinadorIDUseCase *application.GetReportesMantenimientoGeneradoByCoordinadorIDUseCase
}

func NewGetReportesMantenimientoGeneradoByCoordinadorIDController(getByCoordinadorIDUseCase *application.GetReportesMantenimientoGeneradoByCoordinadorIDUseCase) *GetReportesMantenimientoGeneradoByCoordinadorIDController {
	return &GetReportesMantenimientoGeneradoByCoordinadorIDController{
		getByCoordinadorIDUseCase: getByCoordinadorIDUseCase,
	}
}

// @Summary      Reportes por coordinador
// @Tags         ReporteMantenimientoGenerado
// @Produce      json
// @Success      200 {object} entities.RegistroMantenimientoResponse
// @Failure      400 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/reportes-mantenimiento-generado/coordinador/{coordinador_id} [get]
func (ctrl *GetReportesMantenimientoGeneradoByCoordinadorIDController) Run(c *gin.Context) {
	coordinadorIDParam := c.Param("coordinador_id")
	coordinadorID, err := strconv.Atoi(coordinadorIDParam)
	if err != nil {
		core.RespondInvalidInput(c, "ID de coordinador inválido")
		return
	}

	reportes, err := ctrl.getByCoordinadorIDUseCase.Run(int32(coordinadorID))
	if err != nil {
		core.RespondInternalServerError(c, "No se pudieron obtener los reportes del coordinador", err)
		return
	}

	core.RespondOK(c, map[string]interface{}{
		"data": reportes,
	})
}
