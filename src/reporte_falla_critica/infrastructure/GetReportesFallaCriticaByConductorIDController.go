// GetReportesFallaCriticaByConductorIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/reporte_falla_critica/application"
)

type GetReportesFallaCriticaByConductorIDController struct {
	getByConductorIDUseCase *application.GetReportesFallaCriticaByConductorIDUseCase
}

func NewGetReportesFallaCriticaByConductorIDController(getByConductorIDUseCase *application.GetReportesFallaCriticaByConductorIDUseCase) *GetReportesFallaCriticaByConductorIDController {
	return &GetReportesFallaCriticaByConductorIDController{
		getByConductorIDUseCase: getByConductorIDUseCase,
	}
}

func (ctrl *GetReportesFallaCriticaByConductorIDController) Run(c *gin.Context) {
	conductorIDParam := c.Param("conductorId")
	conductorID, err := strconv.Atoi(conductorIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de conductor inválido",
			"error":   err.Error(),
		})
		return
	}

	reportes, err := ctrl.getByConductorIDUseCase.Run(int32(conductorID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los reportes de falla crítica para el conductor",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reportes)
}	