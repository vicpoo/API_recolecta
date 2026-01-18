//CreateIncidenciaController.go
package infrastructure

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/incidencia/application"
	"github.com/vicpoo/API_recolecta/src/incidencia/domain/entities"
)

type CreateIncidenciaController struct {
	createUseCase *application.CreateIncidenciaUseCase
}

func NewCreateIncidenciaController(createUseCase *application.CreateIncidenciaUseCase) *CreateIncidenciaController {
	return &CreateIncidenciaController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateIncidenciaController) Run(c *gin.Context) {
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

	// Parsear fecha
	fechaReporte, err := time.Parse("2006-01-02T15:04:05Z", request.FechaReporte)
	if err != nil {
		// Intentar con otro formato común
		fechaReporte, err = time.Parse("2006-01-02", request.FechaReporte)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Formato de fecha inválido. Use YYYY-MM-DD o YYYY-MM-DDTHH:MM:SSZ",
				"error":   err.Error(),
			})
			return
		}
	}

	// Crear incidencia
	var incidencia *entities.Incidencia
	if request.JsonRuta != "" {
		incidencia = entities.NewIncidenciaConRuta(
			request.PuntoRecoleccionID,
			request.ConductorID,
			request.Descripcion,
			request.JsonRuta,
			fechaReporte,
		)
	} else if request.PuntoRecoleccionID != nil {
		incidencia = entities.NewIncidenciaConPunto(
			request.PuntoRecoleccionID,
			request.ConductorID,
			request.Descripcion,
			fechaReporte,
		)
	} else {
		incidencia = entities.NewIncidencia(
			request.ConductorID,
			request.Descripcion,
			fechaReporte,
		)
	}

	createdIncidencia, err := ctrl.createUseCase.Run(incidencia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear la incidencia",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdIncidencia)
}