//GetNotificacionesGlobalesController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
)

type GetNotificacionesGlobalesController struct {
	useCase *application.GetNotificacionesGlobalesUseCase
}

func NewGetNotificacionesGlobalesController(useCase *application.GetNotificacionesGlobalesUseCase) *GetNotificacionesGlobalesController {
	return &GetNotificacionesGlobalesController{useCase: useCase}
}

func (ctrl *GetNotificacionesGlobalesController) Run(c *gin.Context) {
	result, err := ctrl.useCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudieron obtener las notificaciones globales", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}