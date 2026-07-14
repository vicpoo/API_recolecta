// UpdateRegistroMantenimientoController.go
package infrastructure

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/domain/entities"
	"github.com/vicpoo/API_recolecta/src/core"
)

type UpdateRegistroMantenimientoController struct {
	updateUseCase *application.UpdateRegistroMantenimientoUseCase
}

func NewUpdateRegistroMantenimientoController(updateUseCase *application.UpdateRegistroMantenimientoUseCase) *UpdateRegistroMantenimientoController {
	return &UpdateRegistroMantenimientoController{
		updateUseCase: updateUseCase,
	}
}

// @Summary      Actualizar registro
// @Tags         RegistroMantenimiento
// @Produce      json
// @Param        body body entities.UpdateRegistroMantenimientoRequest true "Body"
// @Param        id path int true "ID"
// @Success      200 {object} entities.RegistroMantenimientoResponse
// @Failure      400 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/registros-mantenimiento/{id} [put]
func (ctrl *UpdateRegistroMantenimientoController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondInvalidInput(c, "ID inválido")
		return
	}

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
	var fechaRealizada time.Time
	if request.FechaRealizada != "" {
		fechaRealizada, err = time.Parse("2006-01-02T15:04:05Z", request.FechaRealizada)
		if err != nil {
			fechaRealizada, err = time.Parse("2006-01-02", request.FechaRealizada)
			if err != nil {
				core.RespondBadRequest(c, "Formato de fecha inválido", map[string]string{"error": err.Error()})
				return
			}
		}
	}

	// Primero obtener el registro existente para mantener algunos campos
	// (En una implementación real, podrías tener un método GetByID en el controller)
	// Por ahora creamos uno nuevo
	registro := &entities.RegistroMantenimiento{
		RegistroID:               int32(id),
		AlertaID:                 request.AlertaID,
		CamionID:                 request.CamionID,
		CoordinadorID:            request.CoordinadorID,
		MecanicoResponsable:      request.MecanicoResponsable,
		KilometrajeMantenimiento: request.KilometrajeMantenimiento,
		Observaciones:            request.Observaciones,
	}

	if !fechaRealizada.IsZero() {
		registro.FechaRealizada = fechaRealizada
	}

	updatedRegistro, err := ctrl.updateUseCase.Run(registro)
	if err != nil {
		if strings.Contains(err.Error(), "no encontrado") {
			core.RespondNotFound(c, "Registro de mantenimiento", idParam)
		} else {
			core.RespondInternalServerError(c, "No se pudo actualizar el registro de mantenimiento", err)
		}
		return
	}

	core.RespondOK(c, updatedRegistro)
}
