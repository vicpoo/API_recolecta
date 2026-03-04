// GetReportesConductorByConductorIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
)

type GetReportesConductorByConductorIDController struct {
	getByConductorIDUseCase *application.GetReportesConductorByConductorIDUseCase
}

func NewGetReportesConductorByConductorIDController(getByConductorIDUseCase *application.GetReportesConductorByConductorIDUseCase) *GetReportesConductorByConductorIDController {
	return &GetReportesConductorByConductorIDController{
		getByConductorIDUseCase: getByConductorIDUseCase,
	}
}

// @Summary      Reportes conductor por conductor
// @Tags         ReporteConductor
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/reportes-conductor/conductor/{conductor_id} [get]
func (ctrl *GetReportesConductorByConductorIDController) Run(c *gin.Context) {
	conductorIDParam := c.Param("conductor_id")
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
			"message": "No se pudieron obtener los reportes del conductor",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reportes)
}