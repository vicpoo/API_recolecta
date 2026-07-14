// GetAllReportesMantenimientoGeneradoController.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetAllReportesMantenimientoGeneradoController struct {
	getAllUseCase *application.GetAllReportesMantenimientoGeneradoUseCase
}

func NewGetAllReportesMantenimientoGeneradoController(getAllUseCase *application.GetAllReportesMantenimientoGeneradoUseCase) *GetAllReportesMantenimientoGeneradoController {
	return &GetAllReportesMantenimientoGeneradoController{
		getAllUseCase: getAllUseCase,
	}
}

// @Summary      Listar reportes generados
// @Tags         ReporteMantenimientoGenerado
// @Produce      json
// @Success      200 {object} entities.RegistroMantenimientoListResponse
// @Failure      400 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/reportes-mantenimiento-generado/ [get]
func (ctrl *GetAllReportesMantenimientoGeneradoController) Run(c *gin.Context) {
	reportes, err := ctrl.getAllUseCase.Run()
	if err != nil {
		core.RespondInternalServerError(c, "No se pudieron obtener los reportes de mantenimiento", err)
		return
	}

	core.RespondOK(c, map[string]interface{}{
		"data": reportes,
	})
}
