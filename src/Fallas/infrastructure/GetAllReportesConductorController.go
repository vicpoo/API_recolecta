// GetAllReportesConductorController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
)

type GetAllReportesConductorController struct {
	getAllUseCase *application.GetAllReportesConductorUseCase
}

func NewGetAllReportesConductorController(getAllUseCase *application.GetAllReportesConductorUseCase) *GetAllReportesConductorController {
	return &GetAllReportesConductorController{
		getAllUseCase: getAllUseCase,
	}
}

// @Summary      Listar reportes conductor
// @Tags         ReporteConductor
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/reportes-conductor/ [get]
func (ctrl *GetAllReportesConductorController) Run(c *gin.Context) {
	reportes, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los reportes",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reportes)
}