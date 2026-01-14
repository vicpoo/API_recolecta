// GetReportesFallaCriticaByCamionIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/reporte_falla_critica/application"
)

type GetReportesFallaCriticaByCamionIDController struct {
	getByCamionIDUseCase *application.GetReportesFallaCriticaByCamionIDUseCase
}

func NewGetReportesFallaCriticaByCamionIDController(getByCamionIDUseCase *application.GetReportesFallaCriticaByCamionIDUseCase) *GetReportesFallaCriticaByCamionIDController {
	return &GetReportesFallaCriticaByCamionIDController{
		getByCamionIDUseCase: getByCamionIDUseCase,
	}
}

func (ctrl *GetReportesFallaCriticaByCamionIDController) Run(c *gin.Context) {
	camionIDParam := c.Param("camionId")
	camionID, err := strconv.Atoi(camionIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de camión inválido",
			"error":   err.Error(),
		})
		return
	}

	reportes, err := ctrl.getByCamionIDUseCase.Run(int32(camionID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los reportes de falla crítica para el camión",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reportes)
}