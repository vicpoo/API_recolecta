// GetRegistrosByFechaRangeController.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetRegistrosByFechaRangeController struct {
	getByFechaRangeUseCase *application.GetRegistrosByFechaRangeUseCase
}

func NewGetRegistrosByFechaRangeController(getByFechaRangeUseCase *application.GetRegistrosByFechaRangeUseCase) *GetRegistrosByFechaRangeController {
	return &GetRegistrosByFechaRangeController{
		getByFechaRangeUseCase: getByFechaRangeUseCase,
	}
}

// @Summary      Registros por rango de fecha
// @Tags         RegistroMantenimiento
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/registros-mantenimiento/fecha [get]
func (ctrl *GetRegistrosByFechaRangeController) Run(c *gin.Context) {
	fechaInicio := c.Query("fecha_inicio")
	fechaFin := c.Query("fecha_fin")

	if fechaInicio == "" || fechaFin == "" {
		core.RespondBadRequest(c, "Se requieren los parámetros fecha_inicio y fecha_fin", map[string]string{
			"fecha_inicio": "requerido",
			"fecha_fin":    "requerido",
		})
		return
	}

	registros, err := ctrl.getByFechaRangeUseCase.Run(fechaInicio, fechaFin)
	if err != nil {
		core.RespondInternalServerError(c, "No se pudieron obtener los registros de mantenimiento", err)
		return
	}

	if len(registros) == 0 {
		core.RespondOK(c, map[string]interface{}{
			"message": "No se encontraron registros de mantenimiento en el rango de fechas especificado",
			"data":    []string{},
		})
		return
	}

	core.RespondOK(c, map[string]interface{}{
		"data": registros,
	})
}
