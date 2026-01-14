//CrearNotificacionFallaController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
)

type CrearNotificacionFallaController struct {
	useCase *application.CrearNotificacionFallaUseCase
}

func NewCrearNotificacionFallaController(useCase *application.CrearNotificacionFallaUseCase) *CrearNotificacionFallaController {
	return &CrearNotificacionFallaController{useCase: useCase}
}

func (ctrl *CrearNotificacionFallaController) Run(c *gin.Context) {
	var request struct {
		UsuarioID  *int32 `json:"usuario_id"`
		Titulo     string `json:"titulo"`
		Mensaje    string `json:"mensaje"`
		CamionID   int32  `json:"camion_id"`
		FallaID    int32  `json:"falla_id"`
		CreadoPor  *int32 `json:"creado_por"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Datos inválidos", "error": err.Error()})
		return
	}

	err := ctrl.useCase.Run(request.UsuarioID, request.Titulo, request.Mensaje, request.CamionID, request.FallaID, request.CreadoPor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudo crear la notificación de falla", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "Notificación de falla creada exitosamente"})
}