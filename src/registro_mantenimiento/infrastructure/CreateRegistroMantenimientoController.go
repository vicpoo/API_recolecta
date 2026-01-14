//CreateRegistroMantenimientoController.go
package infrastructure

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/registro_mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/registro_mantenimiento/domain/entities"
)

type CreateRegistroMantenimientoController struct {
	createUseCase *application.CreateRegistroMantenimientoUseCase
}

func NewCreateRegistroMantenimientoController(createUseCase *application.CreateRegistroMantenimientoUseCase) *CreateRegistroMantenimientoController {
	return &CreateRegistroMantenimientoController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateRegistroMantenimientoController) Run(c *gin.Context) {
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
	fechaRealizada, err := time.Parse("2006-01-02T15:04:05Z", request.FechaRealizada)
	if err != nil {
		// Intentar con otro formato común
		fechaRealizada, err = time.Parse("2006-01-02", request.FechaRealizada)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Formato de fecha inválido. Use YYYY-MM-DD o YYYY-MM-DDTHH:MM:SSZ",
				"error":   err.Error(),
			})
			return
		}
	}

	// Crear registro
	registro := &entities.RegistroMantenimiento{
		AlertaID:                request.AlertaID,
		CamionID:                request.CamionID,
		CoordinadorID:           request.CoordinadorID,
		MecanicoResponsable:     request.MecanicoResponsable,
		FechaRealizada:          fechaRealizada,
		KilometrajeMantenimiento: request.KilometrajeMantenimiento,
		Observaciones:           request.Observaciones,
		CreatedAt:               time.Now(),
	}

	createdRegistro, err := ctrl.createUseCase.Run(registro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear el registro de mantenimiento",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdRegistro)
}