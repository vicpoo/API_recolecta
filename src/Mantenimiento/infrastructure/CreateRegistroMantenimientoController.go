// CreateRegistroMantenimientoController.go
package infrastructure

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/domain/entities"
	"github.com/vicpoo/API_recolecta/src/core"
)

type CreateRegistroMantenimientoController struct {
	createUseCase *application.CreateRegistroMantenimientoUseCase
}

func NewCreateRegistroMantenimientoController(createUseCase *application.CreateRegistroMantenimientoUseCase) *CreateRegistroMantenimientoController {
	return &CreateRegistroMantenimientoController{
		createUseCase: createUseCase,
	}
}

// @Summary      Crear registro de mantenimiento
// @Tags         RegistroMantenimiento
// @Produce      json
// @Param        body body entities.CreateRegistroMantenimientoRequest true "Body"
// @Success      200 {object} entities.RegistroMantenimientoResponse
// @Failure      400 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/registros-mantenimiento/ [post]
func (ctrl *CreateRegistroMantenimientoController) Run(c *gin.Context) {
	var request struct {
		AlertaID                 *int32  `json:"alerta_id"`
		CamionID                 int32   `json:"camion_id"`
		CoordinadorID            int32   `json:"coordinador_id"`
		MecanicoResponsable      string  `json:"mecanico_responsable"`
		FechaRealizada           string  `json:"fecha_realizada"`
		KilometrajeMantenimiento float64 `json:"kilometraje_mantenimiento"`
		Observaciones            string  `json:"observaciones"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		core.RespondBadRequest(c, "Datos inválidos", map[string]string{"error": err.Error()})
		return
	}

	// Parsear fecha
	fechaRealizada, err := time.Parse("2006-01-02T15:04:05Z", request.FechaRealizada)
	if err != nil {
		// Intentar con otro formato común
		fechaRealizada, err = time.Parse("2006-01-02", request.FechaRealizada)
		if err != nil {
			core.RespondBadRequest(c, "Formato de fecha inválido. Use YYYY-MM-DD o YYYY-MM-DDTHH:MM:SSZ", map[string]string{"error": err.Error()})
			return
		}
	}

	// Crear registro
	registro := &entities.RegistroMantenimiento{
		AlertaID:                 request.AlertaID,
		CamionID:                 request.CamionID,
		CoordinadorID:            request.CoordinadorID,
		MecanicoResponsable:      request.MecanicoResponsable,
		FechaRealizada:           fechaRealizada,
		KilometrajeMantenimiento: request.KilometrajeMantenimiento,
		Observaciones:            request.Observaciones,
		CreatedAt:                time.Now(),
	}

	createdRegistro, err := ctrl.createUseCase.Run(registro)
	if err != nil {
		core.RespondInternalServerError(c, "No se pudo crear el registro de mantenimiento", err)
		return
	}

	core.RespondCreated(c, createdRegistro)
}
