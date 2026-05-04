// UpdateAnomaliaController.go
package infrastructure

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/Fallas/domain/entities"
	"github.com/vicpoo/API_recolecta/src/core"
)

type UpdateAnomaliaController struct {
	updateUseCase *application.UpdateAnomaliaUseCase
}

func NewUpdateAnomaliaController(updateUseCase *application.UpdateAnomaliaUseCase) *UpdateAnomaliaController {
	return &UpdateAnomaliaController{
		updateUseCase: updateUseCase,
	}
}

// @Summary      Actualizar anomalía
// @Tags         Anomalia
// @Produce      json
// @Param        body body map[string]interface{} true "Body"
// @Param        id path int true "ID"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/anomalias/{id} [put]
func (ctrl *UpdateAnomaliaController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondInvalidInput(c, "El parámetro 'id' debe ser un número entero válido")
		return
	}

	var request struct {
		PuntoID         *int32  `json:"punto_id"`
		TipoAnomalia    string  `json:"tipo_anomalia" binding:"required"`
		Descripcion     string  `json:"descripcion" binding:"required"`
		FechaReporte    string  `json:"fecha_reporte" binding:"required"`
		Estado          string  `json:"estado" binding:"required"`
		FechaResolucion *string `json:"fecha_resolucion"`
		IDChoferID      int32   `json:"id_chofer_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		core.RespondValidationError(c, "Datos de entrada inválidos", map[string]string{
			"error": err.Error(),
		})
		return
	}

	// Parsear fechas
	fechaReporte, err := parseFecha(request.FechaReporte)
	if err != nil {
		core.RespondBadRequest(c, "Formato de fecha_reporte inválido. Use ISO 8601", map[string]string{
			"field": "fecha_reporte",
			"error": err.Error(),
		})
		return
	}

	var fechaResolucionPtr *time.Time
	if request.FechaResolucion != nil && *request.FechaResolucion != "" {
		fechaResolucion, err := parseFecha(*request.FechaResolucion)
		if err != nil {
			core.RespondBadRequest(c, "Formato de fecha_resolucion inválido. Use ISO 8601", map[string]string{
				"field": "fecha_resolucion",
				"error": err.Error(),
			})
			return
		}
		fechaResolucionPtr = &fechaResolucion
	}

	anomalia := &entities.Anomalia{
		AnomaliaID:      int32(id),
		PuntoID:         request.PuntoID,
		TipoAnomalia:    request.TipoAnomalia,
		Descripcion:     request.Descripcion,
		FechaReporte:    fechaReporte,
		Estado:          request.Estado,
		FechaResolucion: fechaResolucionPtr,
		IDChoferID:      request.IDChoferID,
	}

	updatedAnomalia, err := ctrl.updateUseCase.Run(anomalia)
	if err != nil {
		if err.Error() == "anomalia not found" {
			core.RespondNotFound(c, "Anomalía", idParam)
		} else {
			core.RespondInternalServerError(c, "Error al actualizar la anomalía", err)
		}
		return
	}

	core.RespondOK(c, updatedAnomalia)
}
