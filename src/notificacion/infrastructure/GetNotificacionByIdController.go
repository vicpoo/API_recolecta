//GetNotificacionByIdController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
)

type GetNotificacionByIdController struct {
	getByIdUseCase *application.GetNotificacionByIdUseCase
}

func NewGetNotificacionByIdController(getByIdUseCase *application.GetNotificacionByIdUseCase) *GetNotificacionByIdController {
	return &GetNotificacionByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

func (ctrl *GetNotificacionByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	notificacion, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener la notificación",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, notificacion)
}