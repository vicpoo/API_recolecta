// DeleteAnomaliaController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/anomalia/application"
)

type DeleteAnomaliaController struct {
	deleteUseCase *application.DeleteAnomaliaUseCase
}

func NewDeleteAnomaliaController(deleteUseCase *application.DeleteAnomaliaUseCase) *DeleteAnomaliaController {
	return &DeleteAnomaliaController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteAnomaliaController) Run(c *gin.Context) {
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
			"message": "No se pudo eliminar la anomalía",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Anomalía eliminada exitosamente",
	})
}