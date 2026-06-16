// GetRegistrosByCoordinadorIDController.go
package infrastructure

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetRegistrosByCoordinadorIDController struct {
	getByCoordinadorIDUseCase *application.GetRegistrosByCoordinadorIDUseCase
}

func NewGetRegistrosByCoordinadorIDController(getByCoordinadorIDUseCase *application.GetRegistrosByCoordinadorIDUseCase) *GetRegistrosByCoordinadorIDController {
	return &GetRegistrosByCoordinadorIDController{
		getByCoordinadorIDUseCase: getByCoordinadorIDUseCase,
	}
}

// @Summary      Registros por coordinador
// @Tags         RegistroMantenimiento
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/registros-mantenimiento/coordinador/{coordinador_id} [get]
func (ctrl *GetRegistrosByCoordinadorIDController) Run(c *gin.Context) {
	coordinadorIDParam := c.Param("coordinador_id")
	coordinadorID, err := strconv.Atoi(coordinadorIDParam)
	if err != nil {
		core.RespondInvalidInput(c, "ID de coordinador inválido")
		return
	}

	registros, err := ctrl.getByCoordinadorIDUseCase.Run(int32(coordinadorID))
	if err != nil {
		core.RespondInternalServerError(c, "No se pudieron obtener los registros de mantenimiento del coordinador", err)
		return
	}

	if len(registros) == 0 {
		core.RespondOK(c, map[string]interface{}{
			"message": "No se encontraron registros de mantenimiento para este coordinador",
			"data":    []string{},
		})
		return
	}

	core.RespondOK(c, map[string]interface{}{
		"data": registros,
	})
}
