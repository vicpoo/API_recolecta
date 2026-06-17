// DeleteSeguimientoFallaCriticaController.go
package infrastructure

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type DeleteSeguimientoFallaCriticaController struct {
	deleteUseCase *application.DeleteSeguimientoFallaCriticaUseCase
}

func NewDeleteSeguimientoFallaCriticaController(deleteUseCase *application.DeleteSeguimientoFallaCriticaUseCase) *DeleteSeguimientoFallaCriticaController {
	return &DeleteSeguimientoFallaCriticaController{
		deleteUseCase: deleteUseCase,
	}
}

// @Summary      Eliminar seguimiento falla crítica
// @Tags         SeguimientoFallaCritica
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} entities.SeguimientoFallaCriticaMessageResponse
// @Failure      400 {object} core.ErrorResponse
// @Router       /api/seguimientos-falla-critica/{id} [delete]
func (ctrl *DeleteSeguimientoFallaCriticaController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondBadRequest(c, "ID inválido", map[string]string{"error": err.Error()})
		return
	}

	errDelete := ctrl.deleteUseCase.Run(int32(id))
	if errDelete != nil {
		core.RespondInternalServerError(c, "No se pudo eliminar el seguimiento de falla crítica", errDelete)
		return
	}

	core.RespondOK(c, map[string]string{"status": "Seguimiento de falla crítica eliminado exitosamente"})
}
