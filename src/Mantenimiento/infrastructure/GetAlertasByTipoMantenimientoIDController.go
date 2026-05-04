// GetAlertasByTipoMantenimientoIDController.go
package infrastructure

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetAlertasByTipoMantenimientoIDController struct {
	getByTipoUseCase *application.GetAlertasByTipoMantenimientoIDUseCase
}

func NewGetAlertasByTipoMantenimientoIDController(getByTipoUseCase *application.GetAlertasByTipoMantenimientoIDUseCase) *GetAlertasByTipoMantenimientoIDController {
	return &GetAlertasByTipoMantenimientoIDController{
		getByTipoUseCase: getByTipoUseCase,
	}
}

// @Summary      Alertas por tipo de mantenimiento
// @Tags         AlertaMantenimiento
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/alertas-mantenimiento/tipo/{tipo_id} [get]
func (ctrl *GetAlertasByTipoMantenimientoIDController) Run(c *gin.Context) {
	idParam := c.Param("tipo_id")
	tipoID, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondInvalidInput(c, "El parámetro 'tipo_id' debe ser un número entero válido")
		return
	}

	alertas, err := ctrl.getByTipoUseCase.Run(int32(tipoID))
	if err != nil {
		core.RespondInternalServerError(c, "Error al obtener alertas por tipo de mantenimiento", err)
		return
	}

	core.RespondOK(c, map[string]interface{}{
		"data": alertas,
	})
}
