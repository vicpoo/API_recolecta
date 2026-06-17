// DeleteReporteFallaCriticaController.go
package infrastructure

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type DeleteReporteFallaCriticaController struct {
	deleteUseCase *application.DeleteReporteFallaCriticaUseCase
}

func NewDeleteReporteFallaCriticaController(deleteUseCase *application.DeleteReporteFallaCriticaUseCase) *DeleteReporteFallaCriticaController {
	return &DeleteReporteFallaCriticaController{
		deleteUseCase: deleteUseCase,
	}
}

// @Summary      Eliminar reporte falla crítica
// @Tags         ReporteFallaCritica
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} entities.ReporteFallaCriticaMessageResponse
// @Failure      400 {object} core.ErrorResponse
// @Router       /api/reportes-falla-critica/{id} [delete]
func (ctrl *DeleteReporteFallaCriticaController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondInvalidInput(c, "ID inválido")
		return
	}

	errDelete := ctrl.deleteUseCase.Run(int32(id))
	if errDelete != nil {
		core.RespondInternalServerError(c, "No se pudo eliminar el reporte de falla crítica", errDelete)
		return
	}

	core.RespondOK(c, gin.H{
		"status":  "Reporte de falla crítica marcado como eliminado exitosamente",
		"message": "El reporte ha sido marcado como eliminado (soft delete)",
	})
}
