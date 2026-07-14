package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
)

type NotificarTodosUsuariosController struct {
	useCase *application.NotificarTodosUsuariosUseCase
}

func NewNotificarTodosUsuariosController(useCase *application.NotificarTodosUsuariosUseCase) *NotificarTodosUsuariosController {
	return &NotificarTodosUsuariosController{useCase: useCase}
}

// @Summary      Notificar a todos los usuarios
// @Tags         Notificacion
// @Produce      json
// @Param        body body entities.CreateNotificacionRequest true "Body"
// @Success      200 {object} entities.NotificacionResponse
// @Failure      400 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/notificaciones/enviar-todos [post]
func (ctrl *NotificarTodosUsuariosController) Run(c *gin.Context) {
	var req application.NotificarTodosRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	err := ctrl.useCase.Run(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al enviar notificaciones a todos los usuarios",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Notificaciones enviadas a todos los usuarios exitosamente",
	})
}
