package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
)

type NotificarMultiplesUsuariosController struct {
	useCase *application.NotificarMultiplesUsuariosUseCase
}

func NewNotificarMultiplesUsuariosController(useCase *application.NotificarMultiplesUsuariosUseCase) *NotificarMultiplesUsuariosController {
	return &NotificarMultiplesUsuariosController{useCase: useCase}
}

// @Summary      Notificar a múltiples usuarios
// @Tags         Notificacion
// @Produce      json
// @Param        body body entities.CreateNotificacionRequest true "Body"
// @Success      200 {object} entities.NotificacionResponse
// @Failure      400 {object} core.ErrorResponse
// @Router       /api/notificaciones/enviar-multiples [post]
func (ctrl *NotificarMultiplesUsuariosController) Run(c *gin.Context) {
	var req application.NotificarMultiplesRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	if len(req.UsuarioIDs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Debe especificar al menos un usuario destinatario",
		})
		return
	}

	err := ctrl.useCase.Run(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al enviar notificaciones",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":      "Notificaciones enviadas exitosamente",
		"destinatarios": len(req.UsuarioIDs),
	})
}