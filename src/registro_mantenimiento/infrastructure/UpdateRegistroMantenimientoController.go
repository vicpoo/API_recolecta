//UpdateRegistroMantenimientoController.go
package infrastructure

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/registro_mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/registro_mantenimiento/domain/entities"
)

type UpdateRegistroMantenimientoController struct {
	updateUseCase *application.UpdateRegistroMantenimientoUseCase
}

func NewUpdateRegistroMantenimientoController(updateUseCase *application.UpdateRegistroMantenimientoUseCase) *UpdateRegistroMantenimientoController {
	return &UpdateRegistroMantenimientoController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateRegistroMantenimientoController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	var request struct {
		AlertaID                *int32   `json:"alerta_id"`
		CamionID                int32    `json:"camion_id"`
		CoordinadorID           int32    `json:"coordinador_id"`
		MecanicoResponsable     string   `json:"mecanico_responsable"`
		FechaRealizada          string   `json:"fecha_realizada"`
		KilometrajeMantenimiento float64  `json:"kilometraje_mantenimiento"`
		Observaciones           string   `json:"observaciones"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	// Parsear fecha
	var fechaRealizada time.Time
	if request.FechaRealizada != "" {
		fechaRealizada, err = time.Parse("2006-01-02T15:04:05Z", request.FechaRealizada)
		if err != nil {
			fechaRealizada, err = time.Parse("2006-01-02", request.FechaRealizada)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Formato de fecha inválido",
					"error":   err.Error(),
				})
				return
			}
		}
	}

	// Primero obtener el registro existente para mantener algunos campos
	// (En una implementación real, podrías tener un método GetByID en el controller)
	// Por ahora creamos uno nuevo
	registro := &entities.RegistroMantenimiento{
		RegistroID:              int32(id),
		AlertaID:                request.AlertaID,
		CamionID:                request.CamionID,
		CoordinadorID:           request.CoordinadorID,
		MecanicoResponsable:     request.MecanicoResponsable,
		KilometrajeMantenimiento: request.KilometrajeMantenimiento,
		Observaciones:           request.Observaciones,
	}

	if !fechaRealizada.IsZero() {
		registro.FechaRealizada = fechaRealizada
	}

	updatedRegistro, err := ctrl.updateUseCase.Run(registro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar el registro de mantenimiento",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedRegistro)
}