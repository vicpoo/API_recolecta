// GetAnomaliasByChoferIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/core"
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
// @Success      200 {object} entities.AnomaliaResponse
// @Failure      400 {object} core.ErrorResponse
// @Router       /api/anomalias/chofer/{choferId} [get]
func (ctrl *GetAnomaliasByChoferIDController) Run(c *gin.Context) {
	choferIDParam := c.Param("choferId")
	choferID, err := strconv.Atoi(choferIDParam)
	if err != nil {
		core.RespondInvalidInput(c, "ID de chofer inválido")
		return
	}

	anomalias, err := ctrl.getByChoferIDUseCase.Run(int32(choferID))
	if err != nil {
		core.RespondInternalServerError(c, "No se pudieron obtener las anomalías para el chofer", err)
		return
	}

	c.JSON(http.StatusOK, anomalias)
}
