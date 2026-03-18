// GetReporteConductorByIdController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
)

type GetReporteConductorByIdController struct {
	getByIdUseCase *application.GetReporteConductorByIdUseCase
}

func NewGetReporteConductorByIdController(getByIdUseCase *application.GetReporteConductorByIdUseCase) *GetReporteConductorByIdController {
	return &GetReporteConductorByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

// @Summary      Reporte conductor por ID
// @Tags         ReporteConductor
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/reportes-conductor/{id} [get]
func (ctrl *GetReporteConductorByIdController) Run(c *gin.Context) {
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
			"message": "No se pudo obtener el reporte",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reporte)
}