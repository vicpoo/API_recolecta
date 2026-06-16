// GetReporteFallaCriticaByIdController.go
package infrastructure

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetReporteFallaCriticaByIdController struct {
	getByIdUseCase *application.GetReporteFallaCriticaByIdUseCase
}

func NewGetReporteFallaCriticaByIdController(getByIdUseCase *application.GetReporteFallaCriticaByIdUseCase) *GetReporteFallaCriticaByIdController {
	return &GetReporteFallaCriticaByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

// @Summary      Reporte falla crítica por ID
// @Tags         ReporteFallaCritica
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/reportes-falla-critica/{id} [get]
func (ctrl *GetReporteFallaCriticaByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondInvalidInput(c, "ID inválido")
		return
	}

	reporte, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		core.RespondInternalServerError(c, "No se pudo obtener el reporte de falla crítica", err)
		return
	}

	core.RespondOK(c, reporte)
}
