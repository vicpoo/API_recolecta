//DeleteAlertaMantenimientoController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
)

type DeleteAlertaMantenimientoController struct {
	deleteUseCase *application.DeleteAlertaMantenimientoUseCase
}

func NewDeleteAlertaMantenimientoController(deleteUseCase *application.DeleteAlertaMantenimientoUseCase) *DeleteAlertaMantenimientoController {
	return &DeleteAlertaMantenimientoController{
		deleteUseCase: deleteUseCase,
	}
}

// @Summary      Eliminar alerta
// @Tags         AlertaMantenimiento
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/alertas-mantenimiento/{id} [delete]
func (ctrl *DeleteAlertaMantenimientoController) Run(c *gin.Context) {
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
			"message": "No se pudo eliminar la alerta de mantenimiento",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Alerta de mantenimiento eliminada exitosamente",
	})
}

