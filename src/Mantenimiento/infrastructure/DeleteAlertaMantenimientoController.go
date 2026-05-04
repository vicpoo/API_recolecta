// DeleteAlertaMantenimientoController.go
package infrastructure

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/core"
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
		core.RespondInvalidInput(c, "El parámetro 'id' debe ser un número entero válido")
		return
	}

	errDelete := ctrl.deleteUseCase.Run(int32(id))
	if errDelete != nil {
		if errDelete.Error() == "alerta not found" {
			core.RespondNotFound(c, "Alerta de mantenimiento", idParam)
		} else {
			core.RespondInternalServerError(c, "Error al eliminar la alerta de mantenimiento", errDelete)
		}
		return
	}

	core.RespondOK(c, map[string]string{
		"message": "Alerta de mantenimiento eliminada exitosamente",
	})
}
