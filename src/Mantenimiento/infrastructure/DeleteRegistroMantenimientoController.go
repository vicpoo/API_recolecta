// DeleteRegistroMantenimientoController.go
package infrastructure

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type DeleteRegistroMantenimientoController struct {
	deleteUseCase *application.DeleteRegistroMantenimientoUseCase
}

func NewDeleteRegistroMantenimientoController(deleteUseCase *application.DeleteRegistroMantenimientoUseCase) *DeleteRegistroMantenimientoController {
	return &DeleteRegistroMantenimientoController{
		deleteUseCase: deleteUseCase,
	}
}

// @Summary      Eliminar registro
// @Tags         RegistroMantenimiento
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/registros-mantenimiento/{id} [delete]
func (ctrl *DeleteRegistroMantenimientoController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondInvalidInput(c, "ID inválido")
		return
	}

	errDelete := ctrl.deleteUseCase.Run(int32(id))
	if errDelete != nil {
		if strings.Contains(errDelete.Error(), "no encontrado") {
			core.RespondNotFound(c, "Registro de mantenimiento", idParam)
		} else {
			core.RespondInternalServerError(c, "No se pudo eliminar el registro de mantenimiento", errDelete)
		}
		return
	}

	core.RespondOK(c, map[string]string{
		"status": "Registro de mantenimiento eliminado exitosamente",
	})
}
