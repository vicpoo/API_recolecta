// CreateReporteConductorController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/reporte_conductor/application"
	"github.com/vicpoo/API_recolecta/src/reporte_conductor/domain/entities"
)

type CreateReporteConductorController struct {
	createUseCase *application.CreateReporteConductorUseCase
}

func NewCreateReporteConductorController(createUseCase *application.CreateReporteConductorUseCase) *CreateReporteConductorController {
	return &CreateReporteConductorController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateReporteConductorController) Run(c *gin.Context) {
	var reporteRequest struct {
		ConductorID int32  `json:"conductor_id" binding:"required"`
		CamionID    int32  `json:"camion_id" binding:"required"`
		RutaID      int32  `json:"ruta_id" binding:"required"`
		Descripcion string `json:"descripcion" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reporteRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	reporte := entities.NewReporteConductor(
		reporteRequest.ConductorID,
		reporteRequest.CamionID,
		reporteRequest.RutaID,
		reporteRequest.Descripcion,
	)

	createdReporte, err := ctrl.createUseCase.Run(reporte)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear el reporte del conductor",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdReporte)
}