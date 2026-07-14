// GetReporteConductorByIdController.go
package infrastructure

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetReporteConductorByIdController struct {
	getByIdUseCase *application.GetReporteConductorByIdUseCase
}

func NewGetReporteConductorByIdController(getByIdUseCase *application.GetReporteConductorByIdUseCase) *GetReporteConductorByIdController {
	return &GetReporteConductorByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

// @Summary      Reporte conductor por ID
// @Tags         ReporteConductor
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} entities.ReporteConductorResponse
// @Failure      400 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/reportes-conductor/{id} [get]
func (ctrl *GetReporteConductorByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondBadRequest(c, "ID inválido", map[string]string{"error": err.Error()})
		return
	}

	reporte, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		core.RespondInternalServerError(c, "No se pudo obtener el reporte", err)
		return
	}

	core.RespondOK(c, reporte)
}
