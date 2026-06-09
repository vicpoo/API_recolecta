// UpdateReporteMantenimientoGeneradoController.go
package infrastructure

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/domain/entities"
	"github.com/vicpoo/API_recolecta/src/core"
)

type UpdateReporteMantenimientoGeneradoController struct {
	updateUseCase *application.UpdateReporteMantenimientoGeneradoUseCase
}

func NewUpdateReporteMantenimientoGeneradoController(updateUseCase *application.UpdateReporteMantenimientoGeneradoUseCase) *UpdateReporteMantenimientoGeneradoController {
	return &UpdateReporteMantenimientoGeneradoController{
		updateUseCase: updateUseCase,
	}
}

// @Summary      Actualizar reporte generado
// @Tags         ReporteMantenimientoGenerado
// @Produce      json
// @Param        body body map[string]interface{} true "Body"
// @Param        id path int true "ID"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/reportes-mantenimiento-generado/{id} [put]
func (ctrl *UpdateReporteMantenimientoGeneradoController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondInvalidInput(c, "ID inválido")
		return
	}

	var reporteRequest struct {
		CoordinadorID int32  `json:"coordinador_id"`
		FechaDesde    string `json:"fecha_desde"`
		FechaHasta    string `json:"fecha_hasta"`
		Observaciones string `json:"observaciones"`
	}

	if err := c.ShouldBindJSON(&reporteRequest); err != nil {
		core.RespondBadRequest(c, "Datos inválidos", map[string]string{"error": err.Error()})
		return
	}

	// Parsear fechas si se proporcionan
	var fechaDesde, fechaHasta time.Time
	var errParse error

	if reporteRequest.FechaDesde != "" {
		fechaDesde, errParse = time.Parse(time.RFC3339, reporteRequest.FechaDesde)
		if errParse != nil {
			core.RespondBadRequest(c, "Formato de fecha_desde inválido", map[string]string{"error": errParse.Error()})
			return
		}
	}

	if reporteRequest.FechaHasta != "" {
		fechaHasta, errParse = time.Parse(time.RFC3339, reporteRequest.FechaHasta)
		if errParse != nil {
			core.RespondBadRequest(c, "Formato de fecha_hasta inválido", map[string]string{"error": errParse.Error()})
			return
		}
	}

	reporte := entities.NewReporteMantenimientoGeneradoParaActualizacion(
		int32(id),
		reporteRequest.CoordinadorID,
		fechaDesde,
		fechaHasta,
		reporteRequest.Observaciones,
	)

	updatedReporte, err := ctrl.updateUseCase.Run(reporte)
	if err != nil {
		if strings.Contains(err.Error(), "no encontrado") {
			core.RespondNotFound(c, "Reporte de mantenimiento generado", idParam)
		} else {
			core.RespondInternalServerError(c, "No se pudo actualizar el reporte de mantenimiento", err)
		}
		return
	}

	core.RespondOK(c, updatedReporte)
}
