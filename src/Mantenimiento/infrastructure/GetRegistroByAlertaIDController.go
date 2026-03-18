//GetRegistroByAlertaIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
)

type GetRegistroByAlertaIDController struct {
	getByAlertaIDUseCase *application.GetRegistroByAlertaIDUseCase
}

func NewGetRegistroByAlertaIDController(getByAlertaIDUseCase *application.GetRegistroByAlertaIDUseCase) *GetRegistroByAlertaIDController {
	return &GetRegistroByAlertaIDController{
		getByAlertaIDUseCase: getByAlertaIDUseCase,
	}
}

// @Summary      Registros por alerta
// @Tags         RegistroMantenimiento
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/registros-mantenimiento/alerta/{alerta_id} [get]
func (ctrl *GetRegistroByAlertaIDController) Run(c *gin.Context) {
	alertaIDParam := c.Param("alerta_id")
	alertaID, err := strconv.Atoi(alertaIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de alerta inválido",
			"error":   err.Error(),
		})
		return
	}

	registro, err := ctrl.getByAlertaIDUseCase.Run(int32(alertaID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No se encontró registro de mantenimiento para esta alerta",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, registro)
}