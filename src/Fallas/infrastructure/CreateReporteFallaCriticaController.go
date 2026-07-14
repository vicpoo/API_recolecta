// CreateReporteFallaCriticaController.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/Fallas/domain/entities"
	"github.com/vicpoo/API_recolecta/src/core"
)

type CreateReporteFallaCriticaController struct {
	createUseCase *application.CreateReporteFallaCriticaUseCase
}

func NewCreateReporteFallaCriticaController(createUseCase *application.CreateReporteFallaCriticaUseCase) *CreateReporteFallaCriticaController {
	return &CreateReporteFallaCriticaController{
		createUseCase: createUseCase,
	}
}

// @Summary      Crear reporte falla crítica
// @Tags         ReporteFallaCritica
// @Produce      json
// @Param        body body entities.CreateReporteFallaCriticaRequest true "Body"
// @Success      200 {object} entities.ReporteFallaCriticaResponse
// @Failure      400 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/reportes-falla-critica/ [post]
func (ctrl *CreateReporteFallaCriticaController) Run(c *gin.Context) {
	var request struct {
		CamionID    int32  `json:"camion_id" binding:"required"`
		ConductorID int32  `json:"conductor_id" binding:"required"`
		Descripcion string `json:"descripcion" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		core.RespondBadRequest(c, "Datos inválidos", gin.H{"error": err.Error()})
		return
	}

	reporte := entities.NewReporteFallaCritica(request.CamionID, request.ConductorID, request.Descripcion)

	createdReporte, err := ctrl.createUseCase.Run(reporte)
	if err != nil {
		core.RespondInternalServerError(c, "No se pudo crear el reporte de falla crítica", err)
		return
	}

	core.RespondCreated(c, createdReporte)
}
