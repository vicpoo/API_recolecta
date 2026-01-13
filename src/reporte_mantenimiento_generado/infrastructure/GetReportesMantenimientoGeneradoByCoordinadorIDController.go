// GetReportesMantenimientoGeneradoByCoordinadorIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/reporte_mantenimiento_generado/application"
)

type GetReportesMantenimientoGeneradoByCoordinadorIDController struct {
	getByCoordinadorIDUseCase *application.GetReportesMantenimientoGeneradoByCoordinadorIDUseCase
}

func NewGetReportesMantenimientoGeneradoByCoordinadorIDController(getByCoordinadorIDUseCase *application.GetReportesMantenimientoGeneradoByCoordinadorIDUseCase) *GetReportesMantenimientoGeneradoByCoordinadorIDController {
	return &GetReportesMantenimientoGeneradoByCoordinadorIDController{
		getByCoordinadorIDUseCase: getByCoordinadorIDUseCase,
	}
}

func (ctrl *GetReportesMantenimientoGeneradoByCoordinadorIDController) Run(c *gin.Context) {
	coordinadorIDParam := c.Param("coordinador_id")
	coordinadorID, err := strconv.Atoi(coordinadorIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de coordinador inválido",
			"error":   err.Error(),
		})
		return
	}

	reportes, err := ctrl.getByCoordinadorIDUseCase.Run(int32(coordinadorID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los reportes del coordinador",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reportes)
}