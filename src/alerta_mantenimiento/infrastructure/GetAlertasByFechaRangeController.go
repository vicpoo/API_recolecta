//GetAlertasByFechaRangeController.go
package infrastructure

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/application"
)

type GetAlertasByFechaRangeController struct {
	getByFechaRangeUseCase *application.GetAlertasByFechaRangeUseCase
}

func NewGetAlertasByFechaRangeController(getByFechaRangeUseCase *application.GetAlertasByFechaRangeUseCase) *GetAlertasByFechaRangeController {
	return &GetAlertasByFechaRangeController{
		getByFechaRangeUseCase: getByFechaRangeUseCase,
	}
}

func (ctrl *GetAlertasByFechaRangeController) Run(c *gin.Context) {
	var request struct {
		FechaInicio string `form:"fecha_inicio" binding:"required"`
		FechaFin    string `form:"fecha_fin" binding:"required"`
	}

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Parámetros de fecha inválidos",
			"error":   err.Error(),
		})
		return
	}

	// Parsear fechas
	layout := "2006-01-02"
	fechaInicio, err := time.Parse(layout, request.FechaInicio)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Formato de fecha inicio inválido. Use YYYY-MM-DD",
			"error":   err.Error(),
		})
		return
	}

	fechaFin, err := time.Parse(layout, request.FechaFin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Formato de fecha fin inválido. Use YYYY-MM-DD",
			"error":   err.Error(),
		})
		return
	}

	// Ajustar fecha fin para incluir todo el día
	fechaFin = fechaFin.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	alertas, err := ctrl.getByFechaRangeUseCase.Run(fechaInicio, fechaFin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las alertas en el rango de fechas",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, alertas)
}