// GetAllReportesMantenimientoGeneradoController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/reporte_mantenimiento_generado/application"
)

type GetAllReportesMantenimientoGeneradoController struct {
	getAllUseCase *application.GetAllReportesMantenimientoGeneradoUseCase
}

func NewGetAllReportesMantenimientoGeneradoController(getAllUseCase *application.GetAllReportesMantenimientoGeneradoUseCase) *GetAllReportesMantenimientoGeneradoController {
	return &GetAllReportesMantenimientoGeneradoController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllReportesMantenimientoGeneradoController) Run(c *gin.Context) {
	reportes, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los reportes de mantenimiento",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reportes)
}