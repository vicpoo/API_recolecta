// DeleteReporteMantenimientoGeneradoController.go
package infrastructure

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type DeleteReporteMantenimientoGeneradoController struct {
	deleteUseCase *application.DeleteReporteMantenimientoGeneradoUseCase
}

func NewDeleteReporteMantenimientoGeneradoController(deleteUseCase *application.DeleteReporteMantenimientoGeneradoUseCase) *DeleteReporteMantenimientoGeneradoController {
	return &DeleteReporteMantenimientoGeneradoController{
		deleteUseCase: deleteUseCase,
	}
}

// @Summary      Eliminar reporte generado
// @Tags         ReporteMantenimientoGenerado
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} entities.RegistroMantenimientoMessageResponse
// @Failure      400 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/reportes-mantenimiento-generado/{id} [delete]
func (ctrl *DeleteReporteMantenimientoGeneradoController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondInvalidInput(c, "ID inválido")
		return
	}

	errDelete := ctrl.deleteUseCase.Run(int32(id))
	if errDelete != nil {
		if strings.Contains(errDelete.Error(), "no encontrado") {
			core.RespondNotFound(c, "Reporte de mantenimiento generado", idParam)
		} else {
			core.RespondInternalServerError(c, "No se pudo eliminar el reporte de mantenimiento", errDelete)
		}
		return
	}

	core.RespondOK(c, map[string]string{
		"status": "Reporte de mantenimiento eliminado exitosamente",
	})
}
