// GetReportesConductorByRutaIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/reporte_conductor/application"
)

type GetReportesConductorByRutaIDController struct {
	getByRutaIDUseCase *application.GetReportesConductorByRutaIDUseCase
}

func NewGetReportesConductorByRutaIDController(getByRutaIDUseCase *application.GetReportesConductorByRutaIDUseCase) *GetReportesConductorByRutaIDController {
	return &GetReportesConductorByRutaIDController{
		getByRutaIDUseCase: getByRutaIDUseCase,
	}
}

func (ctrl *GetReportesConductorByRutaIDController) Run(c *gin.Context) {
	rutaIDParam := c.Param("ruta_id")
	rutaID, err := strconv.Atoi(rutaIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de ruta inválido",
			"error":   err.Error(),
		})
		return
	}

	reportes, err := ctrl.getByRutaIDUseCase.Run(int32(rutaID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los reportes de la ruta",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reportes)
}