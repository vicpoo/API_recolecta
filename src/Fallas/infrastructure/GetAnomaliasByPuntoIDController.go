// GetAnomaliasByPuntoIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
)

type GetAnomaliasByPuntoIDController struct {
	getByPuntoIDUseCase *application.GetAnomaliasByPuntoIDUseCase
}

func NewGetAnomaliasByPuntoIDController(getByPuntoIDUseCase *application.GetAnomaliasByPuntoIDUseCase) *GetAnomaliasByPuntoIDController {
	return &GetAnomaliasByPuntoIDController{
		getByPuntoIDUseCase: getByPuntoIDUseCase,
	}
}

// @Summary      Anomalías por punto
// @Tags         Anomalia
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/anomalias/punto/{puntoId} [get]
func (ctrl *GetAnomaliasByPuntoIDController) Run(c *gin.Context) {
	puntoIDParam := c.Param("puntoId")
	puntoID, err := strconv.Atoi(puntoIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de punto inválido",
			"error":   err.Error(),
		})
		return
	}

	anomalias, err := ctrl.getByPuntoIDUseCase.Run(int32(puntoID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las anomalías para el punto",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, anomalias)
}
