//MarcarNotificacionComoLeidaController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
)

type MarcarNotificacionComoLeidaController struct {
	useCase *application.MarcarNotificacionComoLeidaUseCase
}

func NewMarcarNotificacionComoLeidaController(useCase *application.MarcarNotificacionComoLeidaUseCase) *MarcarNotificacionComoLeidaController {
	return &MarcarNotificacionComoLeidaController{useCase: useCase}
}

func (ctrl *MarcarNotificacionComoLeidaController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido", "error": err.Error()})
		return
	}

	err = ctrl.useCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudo marcar la notificación como leída", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Notificación marcada como leída"})
}