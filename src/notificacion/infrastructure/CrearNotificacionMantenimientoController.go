//CrearNotificacionMantenimientoController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
)

type CrearNotificacionMantenimientoController struct {
	useCase *application.CrearNotificacionMantenimientoUseCase
}

func NewCrearNotificacionMantenimientoController(useCase *application.CrearNotificacionMantenimientoUseCase) *CrearNotificacionMantenimientoController {
	return &CrearNotificacionMantenimientoController{useCase: useCase}
}

func (ctrl *CrearNotificacionMantenimientoController) Run(c *gin.Context) {
	var request struct {
		UsuarioID       *int32 `json:"usuario_id"`
		Titulo          string `json:"titulo"`
		Mensaje         string `json:"mensaje"`
		CamionID        int32  `json:"camion_id"`
		MantenimientoID int32  `json:"mantenimiento_id"`
		CreadoPor       *int32 `json:"creado_por"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Datos inválidos", "error": err.Error()})
		return
	}

	err := ctrl.useCase.Run(request.UsuarioID, request.Titulo, request.Mensaje, request.CamionID, request.MantenimientoID, request.CreadoPor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudo crear la notificación de mantenimiento", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "Notificación de mantenimiento creada exitosamente"})
}