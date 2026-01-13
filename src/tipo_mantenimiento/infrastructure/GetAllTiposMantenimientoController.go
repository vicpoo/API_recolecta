// GetAllTiposMantenimientoController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/tipo_mantenimiento/application"
)

type GetAllTiposMantenimientoController struct {
	getAllUseCase *application.GetAllTiposMantenimientoUseCase
}

func NewGetAllTiposMantenimientoController(getAllUseCase *application.GetAllTiposMantenimientoUseCase) *GetAllTiposMantenimientoController {
	return &GetAllTiposMantenimientoController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllTiposMantenimientoController) Run(c *gin.Context) {
	tiposMantenimiento, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los tipos de mantenimiento",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, tiposMantenimiento)
}