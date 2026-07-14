// CreateReporteConductorController.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/Fallas/domain/entities"
	"github.com/vicpoo/API_recolecta/src/core"
)

type CreateReporteConductorController struct {
	createUseCase *application.CreateReporteConductorUseCase
}

func NewCreateReporteConductorController(createUseCase *application.CreateReporteConductorUseCase) *CreateReporteConductorController {
	return &CreateReporteConductorController{
		createUseCase: createUseCase,
	}
}

// @Summary      Crear reporte conductor
// @Tags         ReporteConductor
// @Produce      json
// @Param        body body entities.CreateReporteConductorRequest true "Body"
// @Success      200 {object} entities.ReporteConductorResponse
// @Failure      400 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/reportes-conductor/ [post]
func (ctrl *CreateReporteConductorController) Run(c *gin.Context) {
	var reporteRequest struct {
		ConductorID int32  `json:"conductor_id" binding:"required"`
		CamionID    int32  `json:"camion_id" binding:"required"`
		RutaID      int32  `json:"ruta_id" binding:"required"`
		Descripcion string `json:"descripcion" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reporteRequest); err != nil {
		core.RespondValidationError(c, "Datos inválidos", map[string]string{"error": err.Error()})
		return
	}

	reporte := entities.NewReporteConductor(
		reporteRequest.ConductorID,
		reporteRequest.CamionID,
		reporteRequest.RutaID,
		reporteRequest.Descripcion,
	)

	createdReporte, err := ctrl.createUseCase.Run(reporte)
	if err != nil {
		core.RespondInternalServerError(c, "No se pudo crear el reporte del conductor", err)
		return
	}

	core.RespondCreated(c, createdReporte)
}
