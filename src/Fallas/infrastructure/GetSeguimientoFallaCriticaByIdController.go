// GetSeguimientoFallaCriticaByIdController.go
package infrastructure

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetSeguimientoFallaCriticaByIdController struct {
	getByIdUseCase *application.GetSeguimientoFallaCriticaByIdUseCase
}

func NewGetSeguimientoFallaCriticaByIdController(getByIdUseCase *application.GetSeguimientoFallaCriticaByIdUseCase) *GetSeguimientoFallaCriticaByIdController {
	return &GetSeguimientoFallaCriticaByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

// @Summary      Seguimiento falla crítica por ID
// @Tags         SeguimientoFallaCritica
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} entities.SeguimientoFallaCriticaResponse
// @Failure      400 {object} core.ErrorResponse
// @Router       /api/seguimientos-falla-critica/{id} [get]
func (ctrl *GetSeguimientoFallaCriticaByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondBadRequest(c, "ID inválido", map[string]string{"error": err.Error()})
		return
	}

	seguimiento, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		core.RespondInternalServerError(c, "No se pudo obtener el seguimiento de falla crítica", err)
		return
	}

	core.RespondOK(c, seguimiento)
}
