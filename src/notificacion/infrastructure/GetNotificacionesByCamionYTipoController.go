//GetNotificacionesByCamionYTipoController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
)

type GetNotificacionesByCamionYTipoController struct {
	useCase *application.GetNotificacionesByCamionYTipoUseCase
}

func NewGetNotificacionesByCamionYTipoController(useCase *application.GetNotificacionesByCamionYTipoUseCase) *GetNotificacionesByCamionYTipoController {
	return &GetNotificacionesByCamionYTipoController{useCase: useCase}
}

func (ctrl *GetNotificacionesByCamionYTipoController) Run(c *gin.Context) {
	camionIDParam := c.Param("camion_id")
	camionID, err := strconv.Atoi(camionIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID de camión inválido", "error": err.Error()})
		return
	}

	tipo := c.Param("tipo")
	
	result, err := ctrl.useCase.Run(int32(camionID), tipo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudieron obtener las notificaciones", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}