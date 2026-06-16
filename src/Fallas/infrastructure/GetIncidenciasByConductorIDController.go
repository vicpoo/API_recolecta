// GetIncidenciasByConductorIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetIncidenciasByConductorIDController struct {
	getByConductorIDUseCase *application.GetIncidenciasByConductorIDUseCase
}

func NewGetIncidenciasByConductorIDController(getByConductorIDUseCase *application.GetIncidenciasByConductorIDUseCase) *GetIncidenciasByConductorIDController {
	return &GetIncidenciasByConductorIDController{
		getByConductorIDUseCase: getByConductorIDUseCase,
	}
}

// @Summary      Incidencias por conductor
// @Tags         Incidencia
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/incidencias/conductor/{conductor_id} [get]
func (ctrl *GetIncidenciasByConductorIDController) Run(c *gin.Context) {
	conductorIDParam := c.Param("conductor_id")
	conductorID, err := strconv.Atoi(conductorIDParam)
	if err != nil {
		core.RespondInvalidInput(c, "ID de conductor inválido")
		return
	}

	incidencias, err := ctrl.getByConductorIDUseCase.Run(int32(conductorID))
	if err != nil {
		core.RespondInternalServerError(c, "No se pudieron obtener las incidencias del conductor", err)
		return
	}

	if len(incidencias) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No se encontraron incidencias para este conductor",
			"data":    []string{},
		})
		return
	}

	c.JSON(http.StatusOK, incidencias)
}
