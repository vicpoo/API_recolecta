// GetTipoMantenimientoByIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/tipo_mantenimiento/application"
)

type GetTipoMantenimientoByIDController struct {
	getByIDUseCase *application.GetTipoMantenimientoByIDUseCase
}

func NewGetTipoMantenimientoByIDController(getByIDUseCase *application.GetTipoMantenimientoByIDUseCase) *GetTipoMantenimientoByIDController {
	return &GetTipoMantenimientoByIDController{
		getByIDUseCase: getByIDUseCase,
	}
}

func (ctrl *GetTipoMantenimientoByIDController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	tipoMantenimiento, err := ctrl.getByIDUseCase.Run(int32(id))
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "el tipo de mantenimiento ha sido eliminado" {
			status = http.StatusNotFound
		}
		
		c.JSON(status, gin.H{
			"message": "No se pudo obtener el tipo de mantenimiento",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, tipoMantenimiento)
}