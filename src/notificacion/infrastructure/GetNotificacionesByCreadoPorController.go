//GetNotificacionesByCreadoPorController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
)

type GetNotificacionesByCreadoPorController struct {
	useCase *application.GetNotificacionesByCreadoPorUseCase
}

func NewGetNotificacionesByCreadoPorController(useCase *application.GetNotificacionesByCreadoPorUseCase) *GetNotificacionesByCreadoPorController {
	return &GetNotificacionesByCreadoPorController{useCase: useCase}
}

func (ctrl *GetNotificacionesByCreadoPorController) Run(c *gin.Context) {
	idParam := c.Param("creado_por")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido", "error": err.Error()})
		return
	}

	result, err := ctrl.useCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudieron obtener las notificaciones", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}