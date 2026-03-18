// GetAllReportesFallaCriticaController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
)

type GetAllReportesFallaCriticaController struct {
	getAllUseCase *application.GetAllReportesFallaCriticaUseCase
}

func NewGetAllReportesFallaCriticaController(getAllUseCase *application.GetAllReportesFallaCriticaUseCase) *GetAllReportesFallaCriticaController {
	return &GetAllReportesFallaCriticaController{
		getAllUseCase: getAllUseCase,
	}
}

// @Summary      Listar reportes falla crítica
// @Tags         ReporteFallaCritica
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/reportes-falla-critica/ [get]
func (ctrl *GetAllReportesFallaCriticaController) Run(c *gin.Context) {
	reportes, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los reportes de falla crítica",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reportes)
}