//GetNotificacionesByFallaIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
)

type GetNotificacionesByFallaIDController struct {
	useCase *application.GetNotificacionesByFallaIDUseCase
}

func NewGetNotificacionesByFallaIDController(useCase *application.GetNotificacionesByFallaIDUseCase) *GetNotificacionesByFallaIDController {
	return &GetNotificacionesByFallaIDController{useCase: useCase}
}

func (ctrl *GetNotificacionesByFallaIDController) Run(c *gin.Context) {
	idParam := c.Param("falla_id")
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