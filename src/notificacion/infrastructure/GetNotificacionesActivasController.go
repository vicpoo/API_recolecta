//GetNotificacionesActivasController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
)

type GetNotificacionesActivasController struct {
	useCase *application.GetNotificacionesActivasUseCase
}

func NewGetNotificacionesActivasController(useCase *application.GetNotificacionesActivasUseCase) *GetNotificacionesActivasController {
	return &GetNotificacionesActivasController{useCase: useCase}
}

func (ctrl *GetNotificacionesActivasController) Run(c *gin.Context) {
	result, err := ctrl.useCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudieron obtener las notificaciones activas", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}