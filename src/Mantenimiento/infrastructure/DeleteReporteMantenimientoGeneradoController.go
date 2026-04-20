// DeleteReporteMantenimientoGeneradoController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
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
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/reportes-mantenimiento-generado/{id} [delete]
func (ctrl *DeleteReporteMantenimientoGeneradoController) Run(c *gin.Context) {
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
			"message": "No se pudo eliminar el reporte de mantenimiento",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Reporte de mantenimiento eliminado exitosamente",
	})
}