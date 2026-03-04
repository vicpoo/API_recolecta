//GetAlertasAtendidasController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
)

type GetAlertasAtendidasController struct {
	getAtendidasUseCase *application.GetAlertasAtendidasUseCase
}

func NewGetAlertasAtendidasController(getAtendidasUseCase *application.GetAlertasAtendidasUseCase) *GetAlertasAtendidasController {
	return &GetAlertasAtendidasController{
		getAtendidasUseCase: getAtendidasUseCase,
	}
}

// @Summary      Alertas atendidas
// @Tags         AlertaMantenimiento
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/alertas-mantenimiento/atendidas [get]
func (ctrl *GetAlertasAtendidasController) Run(c *gin.Context) {
	alertas, err := ctrl.getAtendidasUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las alertas atendidas",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, alertas)
}