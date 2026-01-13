//DeleteNotificacionController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
)

type DeleteNotificacionController struct {
	deleteUseCase *application.DeleteNotificacionUseCase
}

func NewDeleteNotificacionController(deleteUseCase *application.DeleteNotificacionUseCase) *DeleteNotificacionController {
	return &DeleteNotificacionController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteNotificacionController) Run(c *gin.Context) {
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
			"message": "No se pudo eliminar la notificación",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Notificación eliminada exitosamente",
	})
}