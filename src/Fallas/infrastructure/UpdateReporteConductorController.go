// UpdateReporteConductorController.go
package infrastructure

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/Fallas/domain/entities"
	"github.com/vicpoo/API_recolecta/src/core"
)

type UpdateReporteConductorController struct {
	updateUseCase *application.UpdateReporteConductorUseCase
}

func NewUpdateReporteConductorController(updateUseCase *application.UpdateReporteConductorUseCase) *UpdateReporteConductorController {
	return &UpdateReporteConductorController{
		updateUseCase: updateUseCase,
	}
}

// @Summary      Actualizar reporte conductor
// @Tags         ReporteConductor
// @Produce      json
// @Param        body body map[string]interface{} true "Body"
// @Param        id path int true "ID"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/reportes-conductor/{id} [put]
func (ctrl *UpdateReporteConductorController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondBadRequest(c, "ID inválido", map[string]string{"error": err.Error()})
		return
	}

	var reporteRequest struct {
		ConductorID int32  `json:"conductor_id"`
		CamionID    int32  `json:"camion_id"`
		RutaID      int32  `json:"ruta_id"`
		Descripcion string `json:"descripcion"`
	}

	if err := c.ShouldBindJSON(&reporteRequest); err != nil {
		core.RespondValidationError(c, "Datos inválidos", map[string]string{"error": err.Error()})
		return
	}

	// Usar el constructor para actualizaciones
	reporte := entities.NewReporteConductorParaActualizacion(
		int32(id),
		reporteRequest.ConductorID,
		reporteRequest.CamionID,
		reporteRequest.RutaID,
		reporteRequest.Descripcion,
	)

	updatedReporte, err := ctrl.updateUseCase.Run(reporte)
	if err != nil {
		core.RespondInternalServerError(c, "No se pudo actualizar el reporte", err)
		return
	}

	core.RespondOK(c, updatedReporte)
}
