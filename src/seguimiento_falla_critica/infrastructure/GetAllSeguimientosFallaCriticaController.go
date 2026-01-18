// GetAllSeguimientosFallaCriticaController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/seguimiento_falla_critica/application"
)

type GetAllSeguimientosFallaCriticaController struct {
	getAllUseCase *application.GetAllSeguimientosFallaCriticaUseCase
}

func NewGetAllSeguimientosFallaCriticaController(getAllUseCase *application.GetAllSeguimientosFallaCriticaUseCase) *GetAllSeguimientosFallaCriticaController {
	return &GetAllSeguimientosFallaCriticaController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllSeguimientosFallaCriticaController) Run(c *gin.Context) {
	seguimientos, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los seguimientos de falla crítica",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, seguimientos)
}