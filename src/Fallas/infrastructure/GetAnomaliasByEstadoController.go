// GetAnomaliasByEstadoController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
)

type GetAnomaliasByEstadoController struct {
	getByEstadoUseCase *application.GetAnomaliasByEstadoUseCase
}

func NewGetAnomaliasByEstadoController(getByEstadoUseCase *application.GetAnomaliasByEstadoUseCase) *GetAnomaliasByEstadoController {
	return &GetAnomaliasByEstadoController{
		getByEstadoUseCase: getByEstadoUseCase,
	}
}

// @Summary      Anomalías por estado
// @Tags         Anomalia
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/anomalias/estado [get]
func (ctrl *GetAnomaliasByEstadoController) Run(c *gin.Context) {
	estado := c.Query("estado")
	if estado == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Se requiere el parámetro estado",
		})
		return
	}

	anomalias, err := ctrl.getByEstadoUseCase.Run(estado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las anomalías por estado",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, anomalias)
}
