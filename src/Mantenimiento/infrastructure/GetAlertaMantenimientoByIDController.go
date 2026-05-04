// GetAlertaMantenimientoByIDController.go
package infrastructure

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetAlertaMantenimientoByIDController struct {
	getByIDUseCase *application.GetAlertaMantenimientoByIDUseCase
}

func NewGetAlertaMantenimientoByIDController(getByIDUseCase *application.GetAlertaMantenimientoByIDUseCase) *GetAlertaMantenimientoByIDController {
	return &GetAlertaMantenimientoByIDController{
		getByIDUseCase: getByIDUseCase,
	}
}

// @Summary      Alerta por ID
// @Tags         AlertaMantenimiento
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/alertas-mantenimiento/{id} [get]
func (ctrl *GetAlertaMantenimientoByIDController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondInvalidInput(c, "El parámetro 'id' debe ser un número entero válido")
		return
	}

	alerta, err := ctrl.getByIDUseCase.Run(int32(id))
	if err != nil {
		if err.Error() == "alerta not found" || alerta == nil {
			core.RespondNotFound(c, "Alerta de mantenimiento", idParam)
		} else {
			core.RespondInternalServerError(c, "Error al obtener la alerta de mantenimiento", err)
		}
		return
	}

	core.RespondOK(c, alerta)
}
