// UpdateReporteMantenimientoGeneradoController.go
package infrastructure

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/reporte_mantenimiento_generado/application"
	"github.com/vicpoo/API_recolecta/src/reporte_mantenimiento_generado/domain/entities"
)

type UpdateReporteMantenimientoGeneradoController struct {
	updateUseCase *application.UpdateReporteMantenimientoGeneradoUseCase
}

func NewUpdateReporteMantenimientoGeneradoController(updateUseCase *application.UpdateReporteMantenimientoGeneradoUseCase) *UpdateReporteMantenimientoGeneradoController {
	return &UpdateReporteMantenimientoGeneradoController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateReporteMantenimientoGeneradoController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	var reporteRequest struct {
		CoordinadorID int32  `json:"coordinador_id"`
		FechaDesde    string `json:"fecha_desde"`
		FechaHasta    string `json:"fecha_hasta"`
		Observaciones string `json:"observaciones"`
	}

	if err := c.ShouldBindJSON(&reporteRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	// Parsear fechas si se proporcionan
	var fechaDesde, fechaHasta time.Time
	var errParse error

	if reporteRequest.FechaDesde != "" {
		fechaDesde, errParse = time.Parse(time.RFC3339, reporteRequest.FechaDesde)
		if errParse != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Formato de fecha_desde inválido",
				"error":   errParse.Error(),
			})
			return
		}
	}

	if reporteRequest.FechaHasta != "" {
		fechaHasta, errParse = time.Parse(time.RFC3339, reporteRequest.FechaHasta)
		if errParse != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Formato de fecha_hasta inválido",
				"error":   errParse.Error(),
			})
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar el reporte de mantenimiento",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedReporte)
}