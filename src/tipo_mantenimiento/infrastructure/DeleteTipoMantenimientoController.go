// DeleteTipoMantenimientoController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/tipo_mantenimiento/application"
)

type DeleteTipoMantenimientoController struct {
	deleteUseCase *application.DeleteTipoMantenimientoUseCase
}

func NewDeleteTipoMantenimientoController(deleteUseCase *application.DeleteTipoMantenimientoUseCase) *DeleteTipoMantenimientoController {
	return &DeleteTipoMantenimientoController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteTipoMantenimientoController) Run(c *gin.Context) {
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
			"message": "No se pudo eliminar el tipo de mantenimiento",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tipo de mantenimiento marcado como eliminado exitosamente",
	})
}