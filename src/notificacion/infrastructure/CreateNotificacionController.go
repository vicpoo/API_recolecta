//CreateNotificacionController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain/entities"
)

type CreateNotificacionController struct {
	createUseCase *application.CreateNotificacionUseCase
}

func NewCreateNotificacionController(createUseCase *application.CreateNotificacionUseCase) *CreateNotificacionController {
	return &CreateNotificacionController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateNotificacionController) Run(c *gin.Context) {
	var request struct {
		UsuarioID                  *int32  `json:"usuario_id"`
		Tipo                       string  `json:"tipo"`
		Titulo                     string  `json:"titulo"`
		Mensaje                    string  `json:"mensaje"`
		IDCamionRelacionado        *int32  `json:"id_camion_relacionado,omitempty"`
		IDFallaRelacionado         *int32  `json:"id_falla_relacionado,omitempty"`
		IDMantenimientoRelacionado *int32  `json:"id_mantenimiento_relacionado,omitempty"`
		CreadoPor                  *int32  `json:"creado_por"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	notificacion := &entities.Notificacion{
		UsuarioID:                  request.UsuarioID,
		Tipo:                       request.Tipo,
		Titulo:                     request.Titulo,
		Mensaje:                    request.Mensaje,
		Activa:                     true,
		IDCamionRelacionado:        request.IDCamionRelacionado,
		IDFallaRelacionado:         request.IDFallaRelacionado,
		IDMantenimientoRelacionado: request.IDMantenimientoRelacionado,
		CreadoPor:                  request.CreadoPor,
	}

	createdNotificacion, err := ctrl.createUseCase.Run(notificacion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear la notificación",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdNotificacion)
}