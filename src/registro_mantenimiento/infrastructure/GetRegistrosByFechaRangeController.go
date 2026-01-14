//GetRegistrosByFechaRangeController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/registro_mantenimiento/application"
)

type GetRegistrosByFechaRangeController struct {
	getByFechaRangeUseCase *application.GetRegistrosByFechaRangeUseCase
}

func NewGetRegistrosByFechaRangeController(getByFechaRangeUseCase *application.GetRegistrosByFechaRangeUseCase) *GetRegistrosByFechaRangeController {
	return &GetRegistrosByFechaRangeController{
		getByFechaRangeUseCase: getByFechaRangeUseCase,
	}
}

func (ctrl *GetRegistrosByFechaRangeController) Run(c *gin.Context) {
	fechaInicio := c.Query("fecha_inicio")
	fechaFin := c.Query("fecha_fin")

	if fechaInicio == "" || fechaFin == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Se requieren los parámetros fecha_inicio y fecha_fin",
		})
		return
	}

	registros, err := ctrl.getByFechaRangeUseCase.Run(fechaInicio, fechaFin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los registros de mantenimiento",
			"error":   err.Error(),
		})
		return
	}

	if len(registros) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No se encontraron registros de mantenimiento en el rango de fechas especificado",
			"data":    []string{},
		})
		return
	}

	c.JSON(http.StatusOK, registros)
}