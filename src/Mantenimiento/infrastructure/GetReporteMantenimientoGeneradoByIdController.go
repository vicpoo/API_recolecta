// GetReporteMantenimientoGeneradoByIdController.go
package infrastructure

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/core"
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
		core.RespondInvalidInput(c, "ID inválido")
		return
	}

	reporte, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		if strings.Contains(err.Error(), "no encontrado") {
			core.RespondNotFound(c, "Reporte de mantenimiento generado", idParam)
		} else {
			core.RespondInternalServerError(c, "No se pudo obtener el reporte de mantenimiento", err)
		}
		return
	}

	core.RespondOK(c, reporte)
}
