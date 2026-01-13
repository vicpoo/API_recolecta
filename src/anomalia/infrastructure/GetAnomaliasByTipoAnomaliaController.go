// GetAnomaliasByTipoAnomaliaController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/anomalia/application"
)

type GetAnomaliasByTipoAnomaliaController struct {
	getByTipoAnomaliaUseCase *application.GetAnomaliasByTipoAnomaliaUseCase
}

func NewGetAnomaliasByTipoAnomaliaController(getByTipoAnomaliaUseCase *application.GetAnomaliasByTipoAnomaliaUseCase) *GetAnomaliasByTipoAnomaliaController {
	return &GetAnomaliasByTipoAnomaliaController{
		getByTipoAnomaliaUseCase: getByTipoAnomaliaUseCase,
	}
}

func (ctrl *GetAnomaliasByTipoAnomaliaController) Run(c *gin.Context) {
	tipoAnomalia := c.Query("tipo_anomalia")
	if tipoAnomalia == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Se requiere el parámetro tipo_anomalia",
		})
		return
	}

	anomalias, err := ctrl.getByTipoAnomaliaUseCase.Run(tipoAnomalia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las anomalías por tipo",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, anomalias)
}