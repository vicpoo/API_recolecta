// UpdateAlertaMantenimientoController.go
package infrastructure

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/domain/entities"
	"github.com/vicpoo/API_recolecta/src/core"
)

type UpdateAlertaMantenimientoController struct {
	updateUseCase *application.UpdateAlertaMantenimientoUseCase
}

func NewUpdateAlertaMantenimientoController(updateUseCase *application.UpdateAlertaMantenimientoUseCase) *UpdateAlertaMantenimientoController {
	return &UpdateAlertaMantenimientoController{
		updateUseCase: updateUseCase,
	}
}

// @Summary      Actualizar alerta
// @Tags         AlertaMantenimiento
// @Produce      json
// @Param        body body map[string]interface{} true "Body"
// @Param        id path int true "ID"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/alertas-mantenimiento/{id} [put]
func (ctrl *UpdateAlertaMantenimientoController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondInvalidInput(c, "El parámetro 'id' debe ser un número entero válido")
		return
	}

	var request struct {
		CamionID            int32  `json:"camion_id" binding:"required"`
		TipoMantenimientoID int32  `json:"tipo_mantenimiento_id" binding:"required"`
		Descripcion         string `json:"descripcion" binding:"required"`
		Observaciones       string `json:"observaciones"`
		Atendido            bool   `json:"atendido"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		core.RespondValidationError(c, "Datos de entrada inválidos", map[string]string{
			"error": err.Error(),
		})
		return
	}

	// Primero obtenemos la alerta existente para preservar la fecha de creación
	alerta, err := ctrl.updateUseCase.Run(&entities.AlertaMantenimiento{
		AlertaID:            int32(id),
		CamionID:            request.CamionID,
		TipoMantenimientoID: request.TipoMantenimientoID,
		Descripcion:         request.Descripcion,
		Observaciones:       request.Observaciones,
		Atendido:            request.Atendido,
	})

	if err != nil {
		if err.Error() == "alerta not found" {
			core.RespondNotFound(c, "Alerta de mantenimiento", idParam)
		} else {
			core.RespondInternalServerError(c, "Error al actualizar la alerta de mantenimiento", err)
		}
		return
	}

	core.RespondOK(c, alerta)
}
