//GetAllAlertasMantenimientoController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
)

type GetAllAlertasMantenimientoController struct {
	getAllUseCase *application.GetAllAlertasMantenimientoUseCase
}

func NewGetAllAlertasMantenimientoController(getAllUseCase *application.GetAllAlertasMantenimientoUseCase) *GetAllAlertasMantenimientoController {
	return &GetAllAlertasMantenimientoController{
		getAllUseCase: getAllUseCase,
	}
}

// @Summary      Listar alertas de mantenimiento
// @Tags         AlertaMantenimiento
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/alertas-mantenimiento/ [get]
func (ctrl *GetAllAlertasMantenimientoController) Run(c *gin.Context) {
	alertas, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las alertas de mantenimiento",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, alertas)
}