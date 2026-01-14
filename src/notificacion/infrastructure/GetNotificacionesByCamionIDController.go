//GetNotificacionesByCamionIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
)

type GetNotificacionesByCamionIDController struct {
	useCase *application.GetNotificacionesByCamionIDUseCase
}

func NewGetNotificacionesByCamionIDController(useCase *application.GetNotificacionesByCamionIDUseCase) *GetNotificacionesByCamionIDController {
	return &GetNotificacionesByCamionIDController{useCase: useCase}
}

func (ctrl *GetNotificacionesByCamionIDController) Run(c *gin.Context) {
	idParam := c.Param("camion_id")
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