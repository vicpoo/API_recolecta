// UpdateReporteConductorController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/reporte_conductor/application"
	"github.com/vicpoo/API_recolecta/src/reporte_conductor/domain/entities"
)

type UpdateReporteConductorController struct {
	updateUseCase *application.UpdateReporteConductorUseCase
}

func NewUpdateReporteConductorController(updateUseCase *application.UpdateReporteConductorUseCase) *UpdateReporteConductorController {
	return &UpdateReporteConductorController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateReporteConductorController) Run(c *gin.Context) {
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
		ConductorID int32  `json:"conductor_id"`
		CamionID    int32  `json:"camion_id"`
		RutaID      int32  `json:"ruta_id"`
		Descripcion string `json:"descripcion"`
	}

	if err := c.ShouldBindJSON(&reporteRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	// Usar el constructor para actualizaciones
	reporte := entities.NewReporteConductorParaActualizacion(
		int32(id),
		reporteRequest.ConductorID,
		reporteRequest.CamionID,
		reporteRequest.RutaID,
		reporteRequest.Descripcion,
	)

	updatedReporte, err := ctrl.updateUseCase.Run(reporte)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar el reporte",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedReporte)
}