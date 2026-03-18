// GetReportesFallaCriticaByConductorIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
)

type GetReportesFallaCriticaByConductorIDController struct {
	getByConductorIDUseCase *application.GetReportesFallaCriticaByConductorIDUseCase
}

func NewGetReportesFallaCriticaByConductorIDController(getByConductorIDUseCase *application.GetReportesFallaCriticaByConductorIDUseCase) *GetReportesFallaCriticaByConductorIDController {
	return &GetReportesFallaCriticaByConductorIDController{
		getByConductorIDUseCase: getByConductorIDUseCase,
	}
}

// @Summary      Reportes falla crítica por conductor
// @Tags         ReporteFallaCritica
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/reportes-falla-critica/conductor/{conductorId} [get]
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