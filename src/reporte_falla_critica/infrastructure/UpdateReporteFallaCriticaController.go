// UpdateReporteFallaCriticaController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/reporte_falla_critica/application"
	"github.com/vicpoo/API_recolecta/src/reporte_falla_critica/domain/entities"
)

type UpdateReporteFallaCriticaController struct {
	updateUseCase *application.UpdateReporteFallaCriticaUseCase
}

func NewUpdateReporteFallaCriticaController(updateUseCase *application.UpdateReporteFallaCriticaUseCase) *UpdateReporteFallaCriticaController {
	return &UpdateReporteFallaCriticaController{
		updateUseCase: updateUseCase,
	}
}

func (ctrl *UpdateReporteFallaCriticaController) Run(c *gin.Context) {
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
		CamionID    int32  `json:"camion_id" binding:"required"`
		ConductorID int32  `json:"conductor_id" binding:"required"`
		Descripcion string `json:"descripcion" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	reporte := &entities.ReporteFallaCritica{
		FallaID:     int32(id),
		CamionID:    request.CamionID,
		ConductorID: request.ConductorID,
		Descripcion: request.Descripcion,
	}

	updatedReporte, err := ctrl.updateUseCase.Run(reporte)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar el reporte de falla crítica",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedReporte)
}