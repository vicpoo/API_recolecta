// GetAllSeguimientosFallaCriticaController.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetAllSeguimientosFallaCriticaController struct {
	getAllUseCase *application.GetAllSeguimientosFallaCriticaUseCase
}

func NewGetAllSeguimientosFallaCriticaController(getAllUseCase *application.GetAllSeguimientosFallaCriticaUseCase) *GetAllSeguimientosFallaCriticaController {
	return &GetAllSeguimientosFallaCriticaController{
		getAllUseCase: getAllUseCase,
	}
}

// @Summary      Listar seguimientos falla crítica
// @Tags         SeguimientoFallaCritica
// @Produce      json
// @Success      200 {object} entities.SeguimientoFallaCriticaListResponse
// @Failure      400 {object} core.ErrorResponse
// @Router       /api/seguimientos-falla-critica/ [get]
func (ctrl *GetAllSeguimientosFallaCriticaController) Run(c *gin.Context) {
	seguimientos, err := ctrl.getAllUseCase.Run()
	if err != nil {
		core.RespondInternalServerError(c, "No se pudieron obtener los seguimientos de falla crítica", err)
		return
	}

	core.RespondOK(c, seguimientos)
}
