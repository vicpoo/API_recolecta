//GetIncidenciasByFechaRangeController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/incidencia/application"
)

type GetIncidenciasByFechaRangeController struct {
	getByFechaRangeUseCase *application.GetIncidenciasByFechaRangeUseCase
}

func NewGetIncidenciasByFechaRangeController(getByFechaRangeUseCase *application.GetIncidenciasByFechaRangeUseCase) *GetIncidenciasByFechaRangeController {
	return &GetIncidenciasByFechaRangeController{
		getByFechaRangeUseCase: getByFechaRangeUseCase,
	}
}

func (ctrl *GetIncidenciasByFechaRangeController) Run(c *gin.Context) {
	fechaInicio := c.Query("fecha_inicio")
	fechaFin := c.Query("fecha_fin")

	if fechaInicio == "" || fechaFin == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Se requieren los parámetros fecha_inicio y fecha_fin",
		})
		return
	}

	incidencias, err := ctrl.getByFechaRangeUseCase.Run(fechaInicio, fechaFin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las incidencias",
			"error":   err.Error(),
		})
		return
	}

	if len(incidencias) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No se encontraron incidencias en el rango de fechas especificado",
			"data":    []string{},
		})
		return
	}

	c.JSON(http.StatusOK, incidencias)
}