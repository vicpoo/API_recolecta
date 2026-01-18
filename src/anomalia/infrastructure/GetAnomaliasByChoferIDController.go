// GetAnomaliasByChoferIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/anomalia/application"
)

type GetAnomaliasByChoferIDController struct {
	getByChoferIDUseCase *application.GetAnomaliasByChoferIDUseCase
}

func NewGetAnomaliasByChoferIDController(getByChoferIDUseCase *application.GetAnomaliasByChoferIDUseCase) *GetAnomaliasByChoferIDController {
	return &GetAnomaliasByChoferIDController{
		getByChoferIDUseCase: getByChoferIDUseCase,
	}
}

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