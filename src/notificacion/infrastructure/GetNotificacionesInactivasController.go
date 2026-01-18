//GetNotificacionesInactivasController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
)

type GetNotificacionesInactivasController struct {
	useCase *application.GetNotificacionesInactivasUseCase
}

func NewGetNotificacionesInactivasController(useCase *application.GetNotificacionesInactivasUseCase) *GetNotificacionesInactivasController {
	return &GetNotificacionesInactivasController{useCase: useCase}
}

func (ctrl *GetNotificacionesInactivasController) Run(c *gin.Context) {
	result, err := ctrl.useCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudieron obtener las notificaciones inactivas", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}