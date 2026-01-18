//NotificarUsuarioController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
)

type NotificarUsuarioController struct {
	useCase *application.NotificarUsuarioUseCase
}

func NewNotificarUsuarioController(useCase *application.NotificarUsuarioUseCase) *NotificarUsuarioController {
	return &NotificarUsuarioController{useCase: useCase}
}

func (ctrl *NotificarUsuarioController) Run(c *gin.Context) {
	var request struct {
		CreadorID           int32  `json:"creador_id"`
		DestinatarioID      int32  `json:"destinatario_id"`
		Tipo                string `json:"tipo"`
		Titulo              string `json:"titulo"`
		Mensaje             string `json:"mensaje"`
		IDCamionRelacionado *int32 `json:"id_camion_relacionado,omitempty"`
		IDFallaRelacionado  *int32 `json:"id_falla_relacionado,omitempty"`
		IDMantenimientoRelacionado *int32 `json:"id_mantenimiento_relacionado,omitempty"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Datos inválidos", "error": err.Error()})
		return
	}

	err := ctrl.useCase.Run(
		request.CreadorID,
		request.DestinatarioID,
		request.Tipo,
		request.Titulo,
		request.Mensaje,
		request.IDCamionRelacionado,
		request.IDFallaRelacionado,
		request.IDMantenimientoRelacionado,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudo enviar la notificación", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "Notificación enviada exitosamente"})
}