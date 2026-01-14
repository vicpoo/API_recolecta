// GetAllReportesConductorController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/reporte_conductor/application"
)

type GetAllReportesConductorController struct {
	getAllUseCase *application.GetAllReportesConductorUseCase
}

func NewGetAllReportesConductorController(getAllUseCase *application.GetAllReportesConductorUseCase) *GetAllReportesConductorController {
	return &GetAllReportesConductorController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllReportesConductorController) Run(c *gin.Context) {
	reportes, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los reportes",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reportes)
}