// GetSeguimientosFallaCriticaByFechaRangeController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/seguimiento_falla_critica/application"
)

type GetSeguimientosFallaCriticaByFechaRangeController struct {
	getByFechaRangeUseCase *application.GetSeguimientosFallaCriticaByFechaRangeUseCase
}

func NewGetSeguimientosFallaCriticaByFechaRangeController(getByFechaRangeUseCase *application.GetSeguimientosFallaCriticaByFechaRangeUseCase) *GetSeguimientosFallaCriticaByFechaRangeController {
	return &GetSeguimientosFallaCriticaByFechaRangeController{
		getByFechaRangeUseCase: getByFechaRangeUseCase,
	}
}

func (ctrl *GetSeguimientosFallaCriticaByFechaRangeController) Run(c *gin.Context) {
	fechaInicio := c.Query("fecha_inicio")
	fechaFin := c.Query("fecha_fin")

	if fechaInicio == "" || fechaFin == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Se requieren los parámetros fecha_inicio y fecha_fin",
		})
		return
	}

	seguimientos, err := ctrl.getByFechaRangeUseCase.Run(fechaInicio, fechaFin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los seguimientos para el rango de fechas",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, seguimientos)
}