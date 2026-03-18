// GetReporteMantenimientoGeneradoByIdController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
)

type GetReporteMantenimientoGeneradoByIdController struct {
	getByIdUseCase *application.GetReporteMantenimientoGeneradoByIdUseCase
}

func NewGetReporteMantenimientoGeneradoByIdController(getByIdUseCase *application.GetReporteMantenimientoGeneradoByIdUseCase) *GetReporteMantenimientoGeneradoByIdController {
	return &GetReporteMantenimientoGeneradoByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

// @Summary      Reporte generado por ID
// @Tags         ReporteMantenimientoGenerado
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/reportes-mantenimiento-generado/{id} [get]
func (ctrl *GetReporteMantenimientoGeneradoByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	reporte, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener el reporte de mantenimiento",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reporte)
}