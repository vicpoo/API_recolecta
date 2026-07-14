// GetAllAlertasMantenimientoController.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetAllAlertasMantenimientoController struct {
	getAllUseCase *application.GetAllAlertasMantenimientoUseCase
}

func NewGetAllAlertasMantenimientoController(getAllUseCase *application.GetAllAlertasMantenimientoUseCase) *GetAllAlertasMantenimientoController {
	return &GetAllAlertasMantenimientoController{
		getAllUseCase: getAllUseCase,
	}
}

// @Summary      Listar alertas de mantenimiento
// @Tags         AlertaMantenimiento
// @Produce      json
// @Success      200 {object} entities.AlertaMantenimientoListResponse
// @Failure      400 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/alertas-mantenimiento/ [get]
func (ctrl *GetAllAlertasMantenimientoController) Run(c *gin.Context) {
	alertas, err := ctrl.getAllUseCase.Run()
	if err != nil {
		core.RespondInternalServerError(c, "Error al obtener las alertas de mantenimiento", err)
		return
	}

	core.RespondOK(c, map[string]interface{}{
		"data": alertas,
	})
}
