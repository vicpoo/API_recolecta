// GetIncidenciasByPuntoRecoleccionIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetIncidenciasByPuntoRecoleccionIDController struct {
	getByPuntoRecoleccionIDUseCase *application.GetIncidenciasByPuntoRecoleccionIDUseCase
}

func NewGetIncidenciasByPuntoRecoleccionIDController(getByPuntoRecoleccionIDUseCase *application.GetIncidenciasByPuntoRecoleccionIDUseCase) *GetIncidenciasByPuntoRecoleccionIDController {
	return &GetIncidenciasByPuntoRecoleccionIDController{
		getByPuntoRecoleccionIDUseCase: getByPuntoRecoleccionIDUseCase,
	}
}

// @Summary      Incidencias por punto
// @Tags         Incidencia
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/incidencias/punto/{punto_recoleccion_id} [get]
func (ctrl *GetIncidenciasByPuntoRecoleccionIDController) Run(c *gin.Context) {
	puntoIDParam := c.Param("punto_recoleccion_id")
	puntoID, err := strconv.Atoi(puntoIDParam)
	if err != nil {
		core.RespondInvalidInput(c, "ID de punto de recolección inválido")
		return
	}

	incidencias, err := ctrl.getByPuntoRecoleccionIDUseCase.Run(int32(puntoID))
	if err != nil {
		core.RespondInternalServerError(c, "No se pudieron obtener las incidencias del punto de recolección", err)
		return
	}

	if len(incidencias) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No se encontraron incidencias para este punto de recolección",
			"data":    []string{},
		})
		return
	}

	c.JSON(http.StatusOK, incidencias)
}
