//UpdateNotificacionController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain/entities"
)

type UpdateNotificacionController struct {
	updateUseCase *application.UpdateNotificacionUseCase
}

func NewUpdateNotificacionController(updateUseCase *application.UpdateNotificacionUseCase) *UpdateNotificacionController {
	return &UpdateNotificacionController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateNotificacionController) Run(c *gin.Context) {
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
		UsuarioID                  *int32  `json:"usuario_id"`
		Tipo                       string  `json:"tipo"`
		Titulo                     string  `json:"titulo"`
		Mensaje                    string  `json:"mensaje"`
		Activa                     bool    `json:"activa"`
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
		NotificacionID:             int32(id),
		UsuarioID:                  request.UsuarioID,
		Tipo:                       request.Tipo,
		Titulo:                     request.Titulo,
		Mensaje:                    request.Mensaje,
		Activa:                     request.Activa,
		IDCamionRelacionado:        request.IDCamionRelacionado,
		IDFallaRelacionado:         request.IDFallaRelacionado,
		IDMantenimientoRelacionado: request.IDMantenimientoRelacionado,
		CreadoPor:                  request.CreadoPor,
	}

	updatedNotificacion, err := ctrl.updateUseCase.Run(notificacion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar la notificación",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedNotificacion)
}