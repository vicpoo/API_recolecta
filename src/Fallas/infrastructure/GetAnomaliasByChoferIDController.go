// GetAnomaliasByChoferIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
)

type GetAnomaliasByChoferIDController struct {
	getByChoferIDUseCase *application.GetAnomaliasByChoferIDUseCase
}

func NewGetAnomaliasByChoferIDController(getByChoferIDUseCase *application.GetAnomaliasByChoferIDUseCase) *GetAnomaliasByChoferIDController {
	return &GetAnomaliasByChoferIDController{
		getByChoferIDUseCase: getByChoferIDUseCase,
	}
}

// @Summary      Anomalías por chofer
// @Tags         Anomalia
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/anomalias/chofer/{choferId} [get]
func (ctrl *GetAnomaliasByChoferIDController) Run(c *gin.Context) {
	choferIDParam := c.Param("choferId")
	choferID, err := strconv.Atoi(choferIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de chofer inválido",
			"error":   err.Error(),
		})
		return
	}

	anomalias, err := ctrl.getByChoferIDUseCase.Run(int32(choferID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las anomalías para el chofer",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, anomalias)
}