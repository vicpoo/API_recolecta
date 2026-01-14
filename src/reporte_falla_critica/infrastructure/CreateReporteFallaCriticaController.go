// CreateReporteFallaCriticaController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/reporte_falla_critica/application"
	"github.com/vicpoo/API_recolecta/src/reporte_falla_critica/domain/entities"
)

type CreateReporteFallaCriticaController struct {
	createUseCase *application.CreateReporteFallaCriticaUseCase
}

func NewCreateReporteFallaCriticaController(createUseCase *application.CreateReporteFallaCriticaUseCase) *CreateReporteFallaCriticaController {
	return &CreateReporteFallaCriticaController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateReporteFallaCriticaController) Run(c *gin.Context) {
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

	reporte := entities.NewReporteFallaCritica(request.CamionID, request.ConductorID, request.Descripcion)

	createdReporte, err := ctrl.createUseCase.Run(reporte)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear el reporte de falla crítica",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdReporte)
}