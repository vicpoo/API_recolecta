//GetIncidenciasByPuntoRecoleccionIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/incidencia/application"
)

type GetIncidenciasByPuntoRecoleccionIDController struct {
	getByPuntoRecoleccionIDUseCase *application.GetIncidenciasByPuntoRecoleccionIDUseCase
}

func NewGetIncidenciasByPuntoRecoleccionIDController(getByPuntoRecoleccionIDUseCase *application.GetIncidenciasByPuntoRecoleccionIDUseCase) *GetIncidenciasByPuntoRecoleccionIDController {
	return &GetIncidenciasByPuntoRecoleccionIDController{
		getByPuntoRecoleccionIDUseCase: getByPuntoRecoleccionIDUseCase,
	}
}

func (ctrl *GetIncidenciasByPuntoRecoleccionIDController) Run(c *gin.Context) {
	puntoIDParam := c.Param("punto_recoleccion_id")
	puntoID, err := strconv.Atoi(puntoIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de punto de recolección inválido",
			"error":   err.Error(),
		})
		return
	}

	incidencias, err := ctrl.getByPuntoRecoleccionIDUseCase.Run(int32(puntoID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las incidencias del punto de recolección",
			"error":   err.Error(),
		})
		return
	}

	if len(incidencias) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No se encontraron incidencias para este punto de recolección",
			"data":    []string{},
		})
		return
	}

	c.JSON(http.StatusOK, incidencias)
}