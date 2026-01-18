// CreateAnomaliaController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/anomalia/application"
	"github.com/vicpoo/API_recolecta/src/anomalia/domain/entities"
)

type CreateAnomaliaController struct {
	createUseCase *application.CreateAnomaliaUseCase
}

func NewCreateAnomaliaController(createUseCase *application.CreateAnomaliaUseCase) *CreateAnomaliaController {
	return &CreateAnomaliaController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateAnomaliaController) Run(c *gin.Context) {
	var request struct {
		PuntoID      *int32  `json:"punto_id"`
		TipoAnomalia string  `json:"tipo_anomalia" binding:"required"`
		Descripcion  string  `json:"descripcion" binding:"required"`
		FechaReporte string  `json:"fecha_reporte" binding:"required"`
		Estado       string  `json:"estado" binding:"required"`
		IDChoferID   int32   `json:"id_chofer_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	// Parsear fecha
	fechaReporte, err := parseFecha(request.FechaReporte)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Formato de fecha inválido. Use YYYY-MM-DD HH:MM:SS",
			"error":   err.Error(),
		})
		return
	}

	anomalia := entities.NewAnomaliaConPunto(
		request.PuntoID,
		request.TipoAnomalia,
		request.Descripcion,
		fechaReporte,
		request.Estado,
		request.IDChoferID,
	)

	createdAnomalia, err := ctrl.createUseCase.Run(anomalia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear la anomalía",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdAnomalia)
}

