//DeleteIncidenciaController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	errDelete := ctrl.deleteUseCase.Run(int32(id))
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo eliminar la incidencia",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Incidencia marcada como eliminada exitosamente",
	})
}