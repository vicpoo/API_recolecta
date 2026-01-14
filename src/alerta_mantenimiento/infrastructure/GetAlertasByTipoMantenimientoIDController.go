//GetAlertasByTipoMantenimientoIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/application"
)

type GetAlertasByTipoMantenimientoIDController struct {
	getByTipoUseCase *application.GetAlertasByTipoMantenimientoIDUseCase
}

func NewGetAlertasByTipoMantenimientoIDController(getByTipoUseCase *application.GetAlertasByTipoMantenimientoIDUseCase) *GetAlertasByTipoMantenimientoIDController {
	return &GetAlertasByTipoMantenimientoIDController{
		getByTipoUseCase: getByTipoUseCase,
	}
}

func (ctrl *GetAlertasByTipoMantenimientoIDController) Run(c *gin.Context) {
	idParam := c.Param("tipo_id")
	tipoID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de tipo de mantenimiento inválido",
			"error":   err.Error(),
		})
		return
	}

	alertas, err := ctrl.getByTipoUseCase.Run(int32(tipoID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las alertas del tipo de mantenimiento",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, alertas)
}