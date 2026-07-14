// GetAllRegistrosMantenimientoController.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetAllRegistrosMantenimientoController struct {
	getAllUseCase *application.GetAllRegistrosMantenimientoUseCase
}

func NewGetAllRegistrosMantenimientoController(getAllUseCase *application.GetAllRegistrosMantenimientoUseCase) *GetAllRegistrosMantenimientoController {
	return &GetAllRegistrosMantenimientoController{
		getAllUseCase: getAllUseCase,
	}
}

// @Summary      Listar registros de mantenimiento
// @Tags         RegistroMantenimiento
// @Produce      json
// @Success      200 {object} entities.RegistroMantenimientoListResponse
// @Failure      400 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/registros-mantenimiento/ [get]
func (ctrl *GetAllRegistrosMantenimientoController) Run(c *gin.Context) {
	registros, err := ctrl.getAllUseCase.Run()
	if err != nil {
		core.RespondInternalServerError(c, "No se pudieron obtener los registros de mantenimiento", err)
		return
	}

	core.RespondOK(c, map[string]interface{}{
		"data": registros,
	})
}
