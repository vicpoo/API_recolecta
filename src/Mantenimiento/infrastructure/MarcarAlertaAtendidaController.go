// MarcarAlertaAtendidaController.go
package infrastructure

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type MarcarAlertaAtendidaController struct {
	marcarAtendidaUseCase *application.MarcarAlertaAtendidaUseCase
}

func NewMarcarAlertaAtendidaController(marcarAtendidaUseCase *application.MarcarAlertaAtendidaUseCase) *MarcarAlertaAtendidaController {
	return &MarcarAlertaAtendidaController{
		marcarAtendidaUseCase: marcarAtendidaUseCase,
	}
}

// @Summary      Marcar alerta como atendida
// @Tags         AlertaMantenimiento
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} entities.AlertaMantenimientoResponse
// @Failure      400 {object} core.ErrorResponse
// @Router       /api/alertas-mantenimiento/{id}/atender [patch]
func (ctrl *MarcarAlertaAtendidaController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondInvalidInput(c, "El parámetro 'id' debe ser un número entero válido")
		return
	}

	err = ctrl.marcarAtendidaUseCase.Run(int32(id))
	if err != nil {
		if err.Error() == "alerta not found" {
			core.RespondNotFound(c, "Alerta de mantenimiento", idParam)
		} else {
			core.RespondInternalServerError(c, "Error al marcar alerta como atendida", err)
		}
		return
	}

	core.RespondOK(c, map[string]string{
		"message": "Alerta marcada como atendida exitosamente",
	})
}
