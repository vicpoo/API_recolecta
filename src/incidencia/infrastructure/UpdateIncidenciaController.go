//UpdateIncidenciaController.go
package infrastructure

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/incidencia/application"
	"github.com/vicpoo/API_recolecta/src/incidencia/domain/entities"
)

type UpdateIncidenciaController struct {
	updateUseCase *application.UpdateIncidenciaUseCase
}

func NewUpdateIncidenciaController(updateUseCase *application.UpdateIncidenciaUseCase) *UpdateIncidenciaController {
	return &UpdateIncidenciaController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateIncidenciaController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	var request struct {
		PuntoRecoleccionID *int32  `json:"punto_recoleccion_id"`
		ConductorID        int32   `json:"conductor_id"`
		Descripcion        string  `json:"descripcion"`
		JsonRuta           string  `json:"json_ruta"`
		FechaReporte       string  `json:"fecha_reporte"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	// Parsear fecha si se proporciona
	var fechaReporte time.Time
	if request.FechaReporte != "" {
		fechaReporte, err = time.Parse("2006-01-02T15:04:05Z", request.FechaReporte)
		if err != nil {
			fechaReporte, err = time.Parse("2006-01-02", request.FechaReporte)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Formato de fecha inválido",
					"error":   err.Error(),
				})
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar la incidencia",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedIncidencia)
}