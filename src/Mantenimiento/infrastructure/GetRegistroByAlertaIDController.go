// GetRegistroByAlertaIDController.go
package infrastructure

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetRegistroByAlertaIDController struct {
	getByAlertaIDUseCase *application.GetRegistroByAlertaIDUseCase
}

func NewGetRegistroByAlertaIDController(getByAlertaIDUseCase *application.GetRegistroByAlertaIDUseCase) *GetRegistroByAlertaIDController {
	return &GetRegistroByAlertaIDController{
		getByAlertaIDUseCase: getByAlertaIDUseCase,
	}
}

// @Summary      Registros por alerta
// @Tags         RegistroMantenimiento
// @Produce      json
// @Success      200 {object} entities.RegistroMantenimientoResponse
// @Failure      400 {object} core.ErrorResponse
// @Router       /api/registros-mantenimiento/alerta/{alerta_id} [get]
func (ctrl *GetRegistroByAlertaIDController) Run(c *gin.Context) {
	alertaIDParam := c.Param("alerta_id")
	alertaID, err := strconv.Atoi(alertaIDParam)
	if err != nil {
		core.RespondInvalidInput(c, "ID de alerta inválido")
		return
	}

	registro, err := ctrl.getByAlertaIDUseCase.Run(int32(alertaID))
	if err != nil {
		if strings.Contains(err.Error(), "no encontrado") {
			core.RespondNotFound(c, "Registro de mantenimiento", alertaIDParam)
		} else {
			core.RespondInternalServerError(c, "No se pudo obtener el registro de mantenimiento para esta alerta", err)
		}
		return
	}

	core.RespondOK(c, registro)
}
