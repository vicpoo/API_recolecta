// CreateReporteMantenimientoGeneradoController.go
package infrastructure

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/domain/entities"
	"github.com/vicpoo/API_recolecta/src/core"
)

type CreateReporteMantenimientoGeneradoController struct {
	createUseCase *application.CreateReporteMantenimientoGeneradoUseCase
}

func NewCreateReporteMantenimientoGeneradoController(createUseCase *application.CreateReporteMantenimientoGeneradoUseCase) *CreateReporteMantenimientoGeneradoController {
	return &CreateReporteMantenimientoGeneradoController{
		createUseCase: createUseCase,
	}
}

// @Summary      Crear reporte de mantenimiento generado
// @Tags         ReporteMantenimientoGenerado
// @Produce      json
// @Param        body body map[string]interface{} true "Body"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/reportes-mantenimiento-generado/ [post]
func (ctrl *CreateReporteMantenimientoGeneradoController) Run(c *gin.Context) {
	var reporteRequest struct {
		CoordinadorID int32  `json:"coordinador_id" binding:"required"`
		FechaDesde    string `json:"fecha_desde" binding:"required"`
		FechaHasta    string `json:"fecha_hasta" binding:"required"`
		Observaciones string `json:"observaciones"`
	}

	if err := c.ShouldBindJSON(&reporteRequest); err != nil {
		core.RespondBadRequest(c, "Datos inválidos", map[string]string{"error": err.Error()})
		return
	}

	// Parsear fechas
	fechaDesde, err := time.Parse(time.RFC3339, reporteRequest.FechaDesde)
	if err != nil {
		core.RespondBadRequest(c, "Formato de fecha_desde inválido. Use RFC3339 (ej: 2024-01-15T00:00:00Z)", map[string]string{"error": err.Error()})
		return
	}

	fechaHasta, err := time.Parse(time.RFC3339, reporteRequest.FechaHasta)
	if err != nil {
		core.RespondBadRequest(c, "Formato de fecha_hasta inválido. Use RFC3339 (ej: 2024-01-15T23:59:59Z)", map[string]string{"error": err.Error()})
		return
	}

	reporte := entities.NewReporteMantenimientoGenerado(
		reporteRequest.CoordinadorID,
		fechaDesde,
		fechaHasta,
		reporteRequest.Observaciones,
	)

	createdReporte, err := ctrl.createUseCase.Run(reporte)
	if err != nil {
		core.RespondInternalServerError(c, "No se pudo crear el reporte de mantenimiento", err)
		return
	}

	core.RespondCreated(c, createdReporte)
}
