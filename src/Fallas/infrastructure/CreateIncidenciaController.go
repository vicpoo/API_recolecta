// CreateIncidenciaController.go
package infrastructure

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/Fallas/domain/entities"
	"github.com/vicpoo/API_recolecta/src/core"
)

type CreateIncidenciaController struct {
	createUseCase *application.CreateIncidenciaUseCase
}

func NewCreateIncidenciaController(createUseCase *application.CreateIncidenciaUseCase) *CreateIncidenciaController {
	return &CreateIncidenciaController{
		createUseCase: createUseCase,
	}
}

// @Summary      Crear incidencia
// @Tags         Incidencia
// @Produce      json
// @Param        body body entities.CreateIncidenciaRequest true "Body"
// @Success      200 {object} entities.IncidenciaResponse
// @Failure      400 {object} core.ErrorResponse
// @Router       /api/incidencias/ [post]
func (ctrl *CreateIncidenciaController) Run(c *gin.Context) {
	var request struct {
		PuntoRecoleccionID *int32 `json:"punto_recoleccion_id"`
		ConductorID        int32  `json:"conductor_id"`
		Descripcion        string `json:"descripcion"`
		JsonRuta           string `json:"json_ruta"`
		FechaReporte       string `json:"fecha_reporte"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		core.RespondValidationError(c, "Datos de incidencia inválidos", map[string]string{"error": err.Error()})
		return
	}

	// Parsear fecha
	fechaReporte, err := time.Parse("2006-01-02T15:04:05Z", request.FechaReporte)
	if err != nil {
		// Intentar con otro formato común
		fechaReporte, err = time.Parse("2006-01-02", request.FechaReporte)
		if err != nil {
			core.RespondValidationError(c, "Formato de fecha inválido", map[string]string{"error": "Use YYYY-MM-DD o YYYY-MM-DDTHH:MM:SSZ"})
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
		core.RespondInternalServerError(c, "Error al crear incidencia", err)
		return
	}

	c.JSON(http.StatusCreated, createdIncidencia)
}
