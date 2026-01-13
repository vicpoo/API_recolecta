// DeleteReporteFallaCriticaController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/reporte_falla_critica/application"
)

type DeleteReporteFallaCriticaController struct {
	deleteUseCase *application.DeleteReporteFallaCriticaUseCase
}

func NewDeleteReporteFallaCriticaController(deleteUseCase *application.DeleteReporteFallaCriticaUseCase) *DeleteReporteFallaCriticaController {
	return &DeleteReporteFallaCriticaController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteReporteFallaCriticaController) Run(c *gin.Context) {
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
			"message": "No se pudo eliminar el reporte de falla crítica",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Reporte de falla crítica marcado como eliminado exitosamente",
		"message": "El reporte ha sido marcado como eliminado (soft delete)",
	})
}