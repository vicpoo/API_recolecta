// GetReporteFallaCriticaByIdController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
)

type GetReporteFallaCriticaByIdController struct {
	getByIdUseCase *application.GetReporteFallaCriticaByIdUseCase
}

func NewGetReporteFallaCriticaByIdController(getByIdUseCase *application.GetReporteFallaCriticaByIdUseCase) *GetReporteFallaCriticaByIdController {
	return &GetReporteFallaCriticaByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

// @Summary      Reporte falla crítica por ID
// @Tags         ReporteFallaCritica
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/reportes-falla-critica/{id} [get]
func (ctrl *GetReporteFallaCriticaByIdController) Run(c *gin.Context) {
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
			"message": "No se pudo obtener el reporte de falla crítica",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reporte)
}