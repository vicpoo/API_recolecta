// GetRegistrosByCamionIDController.go
package infrastructure

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetRegistrosByCamionIDController struct {
	getByCamionIDUseCase *application.GetRegistrosByCamionIDUseCase
}

func NewGetRegistrosByCamionIDController(getByCamionIDUseCase *application.GetRegistrosByCamionIDUseCase) *GetRegistrosByCamionIDController {
	return &GetRegistrosByCamionIDController{
		getByCamionIDUseCase: getByCamionIDUseCase,
	}
}

// @Summary      Registros por camión
// @Tags         RegistroMantenimiento
// @Produce      json
// @Success      200 {object} entities.RegistroMantenimientoListResponse
// @Failure      400 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/registros-mantenimiento/camion/{camion_id} [get]
func (ctrl *GetRegistrosByCamionIDController) Run(c *gin.Context) {
	camionIDParam := c.Param("camion_id")
	camionID, err := strconv.Atoi(camionIDParam)
	if err != nil {
		core.RespondInvalidInput(c, "ID de camión inválido")
		return
	}

	registros, err := ctrl.getByCamionIDUseCase.Run(int32(camionID))
	if err != nil {
		core.RespondInternalServerError(c, "No se pudieron obtener los registros de mantenimiento del camión", err)
		return
	}

	if len(registros) == 0 {
		core.RespondOK(c, map[string]interface{}{
			"message": "No se encontraron registros de mantenimiento para este camión",
			"data":    []string{},
		})
		return
	}

	core.RespondOK(c, map[string]interface{}{
		"data": registros,
	})
}
