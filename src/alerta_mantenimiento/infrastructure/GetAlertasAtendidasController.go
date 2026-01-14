//GetAlertasAtendidasController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/application"
)

type GetAlertasAtendidasController struct {
	getAtendidasUseCase *application.GetAlertasAtendidasUseCase
}

func NewGetAlertasAtendidasController(getAtendidasUseCase *application.GetAlertasAtendidasUseCase) *GetAlertasAtendidasController {
	return &GetAlertasAtendidasController{
		getAtendidasUseCase: getAtendidasUseCase,
	}
}

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