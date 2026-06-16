// GetAlertasPendientesController.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetAlertasPendientesController struct {
	getPendientesUseCase *application.GetAlertasPendientesUseCase
}

func NewGetAlertasPendientesController(getPendientesUseCase *application.GetAlertasPendientesUseCase) *GetAlertasPendientesController {
	return &GetAlertasPendientesController{
		getPendientesUseCase: getPendientesUseCase,
	}
}

// @Summary      Alertas pendientes
// @Tags         AlertaMantenimiento
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/alertas-mantenimiento/pendientes [get]
func (ctrl *GetAlertasPendientesController) Run(c *gin.Context) {
	alertas, err := ctrl.getPendientesUseCase.Run()
	if err != nil {
		core.RespondInternalServerError(c, "Error al obtener alertas pendientes", err)
		return
	}

	core.RespondOK(c, map[string]interface{}{
		"data": alertas,
	})
}
