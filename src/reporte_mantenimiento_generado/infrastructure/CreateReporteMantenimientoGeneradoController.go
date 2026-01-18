// CreateReporteMantenimientoGeneradoController.go
package infrastructure

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/reporte_mantenimiento_generado/application"
	"github.com/vicpoo/API_recolecta/src/reporte_mantenimiento_generado/domain/entities"
)

type CreateReporteMantenimientoGeneradoController struct {
	createUseCase *application.CreateReporteMantenimientoGeneradoUseCase
}

func NewCreateReporteMantenimientoGeneradoController(createUseCase *application.CreateReporteMantenimientoGeneradoUseCase) *CreateReporteMantenimientoGeneradoController {
	return &CreateReporteMantenimientoGeneradoController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateReporteMantenimientoGeneradoController) Run(c *gin.Context) {
	var reporteRequest struct {
		CoordinadorID int32     `json:"coordinador_id" binding:"required"`
		FechaDesde    string    `json:"fecha_desde" binding:"required"`
		FechaHasta    string    `json:"fecha_hasta" binding:"required"`
		Observaciones string    `json:"observaciones"`
	}

	if err := c.ShouldBindJSON(&reporteRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	// Parsear fechas
	fechaDesde, err := time.Parse(time.RFC3339, reporteRequest.FechaDesde)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Formato de fecha_desde inválido. Use RFC3339 (ej: 2024-01-15T00:00:00Z)",
			"error":   err.Error(),
		})
		return
	}

	fechaHasta, err := time.Parse(time.RFC3339, reporteRequest.FechaHasta)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Formato de fecha_hasta inválido. Use RFC3339 (ej: 2024-01-15T23:59:59Z)",
			"error":   err.Error(),
		})
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear el reporte de mantenimiento",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdReporte)
}