// CreateAnomaliaController.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/Fallas/domain/entities"
	"github.com/vicpoo/API_recolecta/src/core"
)

type CreateAnomaliaController struct {
	createUseCase *application.CreateAnomaliaUseCase
}

func NewCreateAnomaliaController(createUseCase *application.CreateAnomaliaUseCase) *CreateAnomaliaController {
	return &CreateAnomaliaController{
		createUseCase: createUseCase,
	}
}

// @Summary      Crear anomalía
// @Tags         Anomalia
// @Produce      json
// @Param        body body entities.CreateAnomaliaRequest true "Body"
// @Success      200 {object} entities.AnomaliaResponse
// @Failure      400 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/anomalias/ [post]
func (ctrl *CreateAnomaliaController) Run(c *gin.Context) {
	var request struct {
		PuntoID      *int32 `json:"punto_id"`
		TipoAnomalia string `json:"tipo_anomalia" binding:"required"`
		Descripcion  string `json:"descripcion" binding:"required"`
		FechaReporte string `json:"fecha_reporte" binding:"required"`
		Estado       string `json:"estado" binding:"required"`
		IDChoferID   int32  `json:"id_chofer_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		core.RespondValidationError(c, "Datos de entrada inválidos", map[string]string{
			"error": err.Error(),
		})
		return
	}

	// Parsear fecha
	fechaReporte, err := parseFecha(request.FechaReporte)
	if err != nil {
		core.RespondBadRequest(c, "Formato de fecha inválido. Use ISO 8601: YYYY-MM-DDTHH:MM:SSZ", map[string]string{
			"field": "fecha_reporte",
			"error": err.Error(),
		})
		return
	}

	anomalia := entities.NewAnomaliaConPunto(
		request.PuntoID,
		request.TipoAnomalia,
		request.Descripcion,
		fechaReporte,
		request.Estado,
		request.IDChoferID,
	)

	createdAnomalia, err := ctrl.createUseCase.Run(anomalia)
	if err != nil {
		core.RespondInternalServerError(c, "No se pudo crear la anomalía", err)
		return
	}

	core.RespondCreated(c, createdAnomalia)
}
