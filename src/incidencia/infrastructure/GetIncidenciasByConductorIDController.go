//GetIncidenciasByConductorIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/incidencia/application"
)

type GetIncidenciasByConductorIDController struct {
	getByConductorIDUseCase *application.GetIncidenciasByConductorIDUseCase
}

func NewGetIncidenciasByConductorIDController(getByConductorIDUseCase *application.GetIncidenciasByConductorIDUseCase) *GetIncidenciasByConductorIDController {
	return &GetIncidenciasByConductorIDController{
		getByConductorIDUseCase: getByConductorIDUseCase,
	}
}

func (ctrl *GetIncidenciasByConductorIDController) Run(c *gin.Context) {
	conductorIDParam := c.Param("conductor_id")
	conductorID, err := strconv.Atoi(conductorIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de conductor inválido",
			"error":   err.Error(),
		})
		return
	}

	incidencias, err := ctrl.getByConductorIDUseCase.Run(int32(conductorID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las incidencias del conductor",
			"error":   err.Error(),
		})
		return
	}

	if len(incidencias) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No se encontraron incidencias para este conductor",
			"data":    []string{},
		})
		return
	}

	c.JSON(http.StatusOK, incidencias)
}