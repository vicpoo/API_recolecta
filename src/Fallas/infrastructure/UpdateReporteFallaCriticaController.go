// UpdateReporteFallaCriticaController.go
package infrastructure

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/Fallas/domain/entities"
	"github.com/vicpoo/API_recolecta/src/core"
)

type UpdateReporteFallaCriticaController struct {
	updateUseCase *application.UpdateReporteFallaCriticaUseCase
}

func NewUpdateReporteFallaCriticaController(updateUseCase *application.UpdateReporteFallaCriticaUseCase) *UpdateReporteFallaCriticaController {
	return &UpdateReporteFallaCriticaController{
		updateUseCase: updateUseCase,
	}
}

// @Summary      Actualizar reporte falla crítica
// @Tags         ReporteFallaCritica
// @Produce      json
// @Param        body body entities.UpdateReporteFallaCriticaRequest true "Body"
// @Param        id path int true "ID"
// @Success      200 {object} entities.ReporteFallaCriticaResponse
// @Failure      400 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/reportes-falla-critica/{id} [put]
func (ctrl *UpdateReporteFallaCriticaController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondInvalidInput(c, "ID inválido")
		return
	}

	var request struct {
		CamionID    int32  `json:"camion_id" binding:"required"`
		ConductorID int32  `json:"conductor_id" binding:"required"`
		Descripcion string `json:"descripcion" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		core.RespondBadRequest(c, "Datos inválidos", gin.H{"error": err.Error()})
		return
	}

	reporte := &entities.ReporteFallaCritica{
		FallaID:     int32(id),
		CamionID:    request.CamionID,
		ConductorID: request.ConductorID,
		Descripcion: request.Descripcion,
	}

	updatedReporte, err := ctrl.updateUseCase.Run(reporte)
	if err != nil {
		core.RespondInternalServerError(c, "No se pudo actualizar el reporte de falla crítica", err)
		return
	}

	core.RespondOK(c, updatedReporte)
}
