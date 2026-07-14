// GetAnomaliasByTipoAnomaliaController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetAnomaliasByTipoAnomaliaController struct {
	getByTipoAnomaliaUseCase *application.GetAnomaliasByTipoAnomaliaUseCase
}

func NewGetAnomaliasByTipoAnomaliaController(getByTipoAnomaliaUseCase *application.GetAnomaliasByTipoAnomaliaUseCase) *GetAnomaliasByTipoAnomaliaController {
	return &GetAnomaliasByTipoAnomaliaController{
		getByTipoAnomaliaUseCase: getByTipoAnomaliaUseCase,
	}
}

// @Summary      Anomalías por tipo
// @Tags         Anomalia
// @Produce      json
// @Success      200 {object} entities.AnomaliaResponse
// @Failure      400 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/anomalias/tipo [get]
func (ctrl *GetAnomaliasByTipoAnomaliaController) Run(c *gin.Context) {
	tipoAnomalia := c.Query("tipo_anomalia")
	if tipoAnomalia == "" {
		core.RespondValidationError(c, "Parámetro de tipo_anomalia inválido", map[string]string{"tipo_anomalia": "requerido"})
		return
	}

	anomalias, err := ctrl.getByTipoAnomaliaUseCase.Run(tipoAnomalia)
	if err != nil {
		core.RespondInternalServerError(c, "No se pudieron obtener las anomalías por tipo", err)
		return
	}

	c.JSON(http.StatusOK, anomalias)
}
