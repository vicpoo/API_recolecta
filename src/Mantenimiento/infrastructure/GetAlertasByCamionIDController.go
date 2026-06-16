// GetAlertasByCamionIDController.go
package infrastructure

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetAlertasByCamionIDController struct {
	getByCamionUseCase *application.GetAlertasByCamionIDUseCase
}

func NewGetAlertasByCamionIDController(getByCamionUseCase *application.GetAlertasByCamionIDUseCase) *GetAlertasByCamionIDController {
	return &GetAlertasByCamionIDController{
		getByCamionUseCase: getByCamionUseCase,
	}
}

// @Summary      Alertas por camión
// @Tags         AlertaMantenimiento
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/alertas-mantenimiento/camion/{camion_id} [get]
func (ctrl *GetAlertasByCamionIDController) Run(c *gin.Context) {
	idParam := c.Param("camion_id")
	camionID, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondInvalidInput(c, "El parámetro 'camion_id' debe ser un número entero válido")
		return
	}

	alertas, err := ctrl.getByCamionUseCase.Run(int32(camionID))
	if err != nil {
		core.RespondInternalServerError(c, "Error al obtener alertas del camión", err)
		return
	}

	core.RespondOK(c, map[string]interface{}{
		"data": alertas,
	})
}
