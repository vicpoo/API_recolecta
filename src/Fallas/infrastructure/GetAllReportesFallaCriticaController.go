// GetAllReportesFallaCriticaController.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetAllReportesFallaCriticaController struct {
	getAllUseCase *application.GetAllReportesFallaCriticaUseCase
}

func NewGetAllReportesFallaCriticaController(getAllUseCase *application.GetAllReportesFallaCriticaUseCase) *GetAllReportesFallaCriticaController {
	return &GetAllReportesFallaCriticaController{
		getAllUseCase: getAllUseCase,
	}
}

// @Summary      Listar reportes falla crítica
// @Tags         ReporteFallaCritica
// @Produce      json
// @Success      200 {object} entities.ReporteFallaCriticaListResponse
// @Failure      400 {object} core.ErrorResponse
// @Router       /api/reportes-falla-critica/ [get]
func (ctrl *GetAllReportesFallaCriticaController) Run(c *gin.Context) {
	reportes, err := ctrl.getAllUseCase.Run()
	if err != nil {
		core.RespondInternalServerError(c, "No se pudieron obtener los reportes de falla crítica", err)
		return
	}

	core.RespondOK(c, reportes)
}
