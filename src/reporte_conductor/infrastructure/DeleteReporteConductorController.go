// DeleteReporteConductorController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/reporte_conductor/application"
)

type DeleteReporteConductorController struct {
	deleteUseCase *application.DeleteReporteConductorUseCase
}

func NewDeleteReporteConductorController(deleteUseCase *application.DeleteReporteConductorUseCase) *DeleteReporteConductorController {
	return &DeleteReporteConductorController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteReporteConductorController) Run(c *gin.Context) {
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
			"message": "No se pudo eliminar el reporte del conductor",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Reporte eliminado exitosamente",
	})
}