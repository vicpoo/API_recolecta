//DeleteRegistroMantenimientoController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/registro_mantenimiento/application"
)

type DeleteRegistroMantenimientoController struct {
	deleteUseCase *application.DeleteRegistroMantenimientoUseCase
}

func NewDeleteRegistroMantenimientoController(deleteUseCase *application.DeleteRegistroMantenimientoUseCase) *DeleteRegistroMantenimientoController {
	return &DeleteRegistroMantenimientoController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteRegistroMantenimientoController) Run(c *gin.Context) {
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
			"message": "No se pudo eliminar el registro de mantenimiento",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Registro de mantenimiento eliminado exitosamente",
	})
}