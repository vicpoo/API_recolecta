// UpdateAnomaliaController.go
package infrastructure

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/anomalia/application"
	"github.com/vicpoo/API_recolecta/src/anomalia/domain/entities"
)

type UpdateAnomaliaController struct {
	updateUseCase *application.UpdateAnomaliaUseCase
}

func NewUpdateAnomaliaController(updateUseCase *application.UpdateAnomaliaUseCase) *UpdateAnomaliaController {
	return &UpdateAnomaliaController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateAnomaliaController) Run(c *gin.Context) {
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
		PuntoID         *int32  `json:"punto_id"`
		TipoAnomalia    string  `json:"tipo_anomalia" binding:"required"`
		Descripcion     string  `json:"descripcion" binding:"required"`
		FechaReporte    string  `json:"fecha_reporte" binding:"required"`
		Estado          string  `json:"estado" binding:"required"`
		FechaResolucion *string `json:"fecha_resolucion"`
		IDChoferID      int32   `json:"id_chofer_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	// Parsear fechas
	fechaReporte, err := parseFecha(request.FechaReporte)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Formato de fecha_reporte inválido",
			"error":   err.Error(),
		})
		return
	}

	var fechaResolucionPtr *time.Time
	if request.FechaResolucion != nil && *request.FechaResolucion != "" {
		fechaResolucion, err := parseFecha(*request.FechaResolucion)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Formato de fecha_resolucion inválido",
				"error":   err.Error(),
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar la anomalía",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedAnomalia)
}
