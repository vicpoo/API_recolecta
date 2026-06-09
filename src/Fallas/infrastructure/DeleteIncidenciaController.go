// DeleteIncidenciaController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type DeleteIncidenciaController struct {
	deleteUseCase *application.DeleteIncidenciaUseCase
}

func NewDeleteIncidenciaController(deleteUseCase *application.DeleteIncidenciaUseCase) *DeleteIncidenciaController {
	return &DeleteIncidenciaController{
		deleteUseCase: deleteUseCase,
	}
}

// @Summary      Eliminar incidencia
// @Tags         Incidencia
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/incidencias/{id} [delete]
func (ctrl *DeleteIncidenciaController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondInvalidInput(c, "ID de incidencia inválido")
		return
	}

	errDelete := ctrl.deleteUseCase.Run(int32(id))
	if errDelete != nil {
		core.RespondInternalServerError(c, "Error al eliminar incidencia", errDelete)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Incidencia marcada como eliminada exitosamente",
	})
}
