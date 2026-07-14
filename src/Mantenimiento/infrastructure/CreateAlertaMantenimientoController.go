// CreateAlertaMantenimientoController.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/domain/entities"
	"github.com/vicpoo/API_recolecta/src/core"
)

type CreateAlertaMantenimientoController struct {
	createUseCase *application.CreateAlertaMantenimientoUseCase
}

func NewCreateAlertaMantenimientoController(createUseCase *application.CreateAlertaMantenimientoUseCase) *CreateAlertaMantenimientoController {
	return &CreateAlertaMantenimientoController{
		createUseCase: createUseCase,
	}
}

// @Summary      Crear alerta de mantenimiento
// @Tags         AlertaMantenimiento
// @Produce      json
// @Param        body body entities.CreateAlertaMantenimientoRequest true "Body"
// @Success      200 {object} entities.AlertaMantenimientoResponse
// @Failure      400 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/alertas-mantenimiento/ [post]
func (ctrl *CreateAlertaMantenimientoController) Run(c *gin.Context) {
	var request struct {
		CamionID            int32  `json:"camion_id" binding:"required"`
		TipoMantenimientoID int32  `json:"tipo_mantenimiento_id" binding:"required"`
		Descripcion         string `json:"descripcion" binding:"required"`
		Observaciones       string `json:"observaciones"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		core.RespondValidationError(c, "Datos de entrada inválidos", map[string]string{
			"error": err.Error(),
		})
		return
	}

	alerta := entities.NewAlertaMantenimiento(
		request.CamionID,
		request.TipoMantenimientoID,
		request.Descripcion,
		request.Observaciones,
	)

	createdAlerta, err := ctrl.createUseCase.Run(alerta)
	if err != nil {
		core.RespondInternalServerError(c, "Error al crear la alerta de mantenimiento", err)
		return
	}

	core.RespondCreated(c, createdAlerta)
}
