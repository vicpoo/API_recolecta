// GetReportesConductorByCamionIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
)

type GetReportesConductorByCamionIDController struct {
	getByCamionIDUseCase *application.GetReportesConductorByCamionIDUseCase
}

func NewGetReportesConductorByCamionIDController(getByCamionIDUseCase *application.GetReportesConductorByCamionIDUseCase) *GetReportesConductorByCamionIDController {
	return &GetReportesConductorByCamionIDController{
		getByCamionIDUseCase: getByCamionIDUseCase,
	}
}

// @Summary      Reportes conductor por camión
// @Tags         ReporteConductor
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/reportes-conductor/camion/{camion_id} [get]
func (ctrl *GetReportesConductorByCamionIDController) Run(c *gin.Context) {
	camionIDParam := c.Param("camion_id")
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
			"message": "No se pudieron obtener los reportes del camión",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reportes)
}