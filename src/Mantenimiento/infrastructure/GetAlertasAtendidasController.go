// GetAlertasAtendidasController.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/core"
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
		core.RespondInternalServerError(c, "Error al obtener alertas atendidas", err)
		return
	}

	core.RespondOK(c, map[string]interface{}{
		"data": alertas,
	})
}
