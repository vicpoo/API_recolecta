//GetAlertaMantenimientoByIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/application"
)

type GetAlertaMantenimientoByIDController struct {
	getByIDUseCase *application.GetAlertaMantenimientoByIDUseCase
}

func NewGetAlertaMantenimientoByIDController(getByIDUseCase *application.GetAlertaMantenimientoByIDUseCase) *GetAlertaMantenimientoByIDController {
	return &GetAlertaMantenimientoByIDController{
		getByIDUseCase: getByIDUseCase,
	}
}

func (ctrl *GetAlertaMantenimientoByIDController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	alerta, err := ctrl.getByIDUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener la alerta de mantenimiento",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, alerta)
}