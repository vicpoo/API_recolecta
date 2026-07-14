// DeleteReporteConductorController.go
package infrastructure

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type DeleteReporteConductorController struct {
	deleteUseCase *application.DeleteReporteConductorUseCase
}

func NewDeleteReporteConductorController(deleteUseCase *application.DeleteReporteConductorUseCase) *DeleteReporteConductorController {
	return &DeleteReporteConductorController{
		deleteUseCase: deleteUseCase,
	}
}

// @Summary      Eliminar reporte conductor
// @Tags         ReporteConductor
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} entities.ReporteConductorMessageResponse
// @Failure      400 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/reportes-conductor/{id} [delete]
func (ctrl *DeleteReporteConductorController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondBadRequest(c, "ID inválido", map[string]string{"error": err.Error()})
		return
	}

	errDelete := ctrl.deleteUseCase.Run(int32(id))
	if errDelete != nil {
		core.RespondInternalServerError(c, "No se pudo eliminar el reporte del conductor", errDelete)
		return
	}

	core.RespondOK(c, gin.H{"status": "Reporte eliminado exitosamente"})
}
