// DeleteSeguimientoFallaCriticaController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
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
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/seguimientos-falla-critica/{id} [delete]
func (ctrl *DeleteSeguimientoFallaCriticaController) Run(c *gin.Context) {
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
			"message": "No se pudo eliminar el seguimiento de falla crítica",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Seguimiento de falla crítica eliminado exitosamente",
	})
}