// UpdateIncidenciaController.go
package infrastructure

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/Fallas/domain/entities"
	"github.com/vicpoo/API_recolecta/src/core"
)

type UpdateIncidenciaController struct {
	updateUseCase *application.UpdateIncidenciaUseCase
}

func NewUpdateIncidenciaController(updateUseCase *application.UpdateIncidenciaUseCase) *UpdateIncidenciaController {
	return &UpdateIncidenciaController{
		updateUseCase: updateUseCase,
	}
}

// @Summary      Actualizar incidencia
// @Tags         Incidencia
// @Produce      json
// @Param        body body entities.UpdateIncidenciaRequest true "Body"
// @Param        id path int true "ID"
// @Success      200 {object} entities.IncidenciaResponse
// @Failure      400 {object} core.ErrorResponse
// @Router       /api/incidencias/{id} [put]
func (ctrl *UpdateIncidenciaController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondInvalidInput(c, "ID de incidencia inválido")
		return
	}

	var request struct {
		PuntoRecoleccionID *int32 `json:"punto_recoleccion_id"`
		ConductorID        int32  `json:"conductor_id"`
		Descripcion        string `json:"descripcion"`
		JsonRuta           string `json:"json_ruta"`
		FechaReporte       string `json:"fecha_reporte"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		core.RespondValidationError(c, "Datos de incidencia inválidos", map[string]string{"error": err.Error()})
		return
	}

	// Parsear fecha si se proporciona
	var fechaReporte time.Time
	if request.FechaReporte != "" {
		fechaReporte, err = time.Parse("2006-01-02T15:04:05Z", request.FechaReporte)
		if err != nil {
			fechaReporte, err = time.Parse("2006-01-02", request.FechaReporte)
			if err != nil {
				core.RespondValidationError(c, "Formato de fecha inválido", map[string]string{"error": err.Error()})
				return
			}
		}
	}

	// Crear objeto incidencia para actualizar
	incidencia := &entities.Incidencia{
		IncidenciaID:       int32(id),
		PuntoRecoleccionID: request.PuntoRecoleccionID,
		ConductorID:        request.ConductorID,
		Descripcion:        request.Descripcion,
		JsonRuta:           request.JsonRuta,
	}

	if !fechaReporte.IsZero() {
		incidencia.FechaReporte = fechaReporte
	}

	updatedIncidencia, err := ctrl.updateUseCase.Run(incidencia)
	if err != nil {
		core.RespondInternalServerError(c, "Error al actualizar incidencia", err)
		return
	}

	c.JSON(http.StatusOK, updatedIncidencia)
}
