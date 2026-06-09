// GetSeguimientosFallaCriticaByFallaIDController.go
package infrastructure

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetSeguimientosFallaCriticaByFallaIDController struct {
	getByFallaIDUseCase *application.GetSeguimientosFallaCriticaByFallaIDUseCase
}

func NewGetSeguimientosFallaCriticaByFallaIDController(getByFallaIDUseCase *application.GetSeguimientosFallaCriticaByFallaIDUseCase) *GetSeguimientosFallaCriticaByFallaIDController {
	return &GetSeguimientosFallaCriticaByFallaIDController{
		getByFallaIDUseCase: getByFallaIDUseCase,
	}
}

// @Summary      Seguimientos por falla
// @Tags         SeguimientoFallaCritica
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/seguimientos-falla-critica/falla/{fallaId} [get]
func (ctrl *GetSeguimientosFallaCriticaByFallaIDController) Run(c *gin.Context) {
	fallaIDParam := c.Param("fallaId")
	fallaID, err := strconv.Atoi(fallaIDParam)
	if err != nil {
		core.RespondBadRequest(c, "ID de falla inválido", map[string]string{"error": err.Error()})
		return
	}

	seguimientos, err := ctrl.getByFallaIDUseCase.Run(int32(fallaID))
	if err != nil {
		core.RespondInternalServerError(c, "No se pudieron obtener los seguimientos para la falla", err)
		return
	}

	core.RespondOK(c, seguimientos)
}
