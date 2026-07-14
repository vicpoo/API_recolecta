// GetIncidenciasByFechaRangeController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetIncidenciasByFechaRangeController struct {
	getByFechaRangeUseCase *application.GetIncidenciasByFechaRangeUseCase
}

func NewGetIncidenciasByFechaRangeController(getByFechaRangeUseCase *application.GetIncidenciasByFechaRangeUseCase) *GetIncidenciasByFechaRangeController {
	return &GetIncidenciasByFechaRangeController{
		getByFechaRangeUseCase: getByFechaRangeUseCase,
	}
}

// @Summary      Incidencias por fecha
// @Tags         Incidencia
// @Produce      json
// @Success      200 {object} entities.IncidenciaListResponse
// @Failure      400 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/incidencias/fecha [get]
func (ctrl *GetIncidenciasByFechaRangeController) Run(c *gin.Context) {
	fechaInicio := c.Query("fecha_inicio")
	fechaFin := c.Query("fecha_fin")

	if fechaInicio == "" || fechaFin == "" {
		core.RespondValidationError(c, "Parámetros de fecha inválidos", map[string]string{
			"fecha_inicio": "requerido",
			"fecha_fin":    "requerido",
		})
		return
	}

	incidencias, err := ctrl.getByFechaRangeUseCase.Run(fechaInicio, fechaFin)
	if err != nil {
		core.RespondInternalServerError(c, "No se pudieron obtener las incidencias", err)
		return
	}

	if len(incidencias) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No se encontraron incidencias en el rango de fechas especificado",
			"data":    []string{},
		})
		return
	}

	c.JSON(http.StatusOK, incidencias)
}
