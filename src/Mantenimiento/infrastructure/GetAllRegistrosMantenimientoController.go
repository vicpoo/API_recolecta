//GetAllRegistrosMantenimientoController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
)

type GetAllRegistrosMantenimientoController struct {
	getAllUseCase *application.GetAllRegistrosMantenimientoUseCase
}

func NewGetAllRegistrosMantenimientoController(getAllUseCase *application.GetAllRegistrosMantenimientoUseCase) *GetAllRegistrosMantenimientoController {
	return &GetAllRegistrosMantenimientoController{
		getAllUseCase: getAllUseCase,
	}
}

// @Summary      Listar registros de mantenimiento
// @Tags         RegistroMantenimiento
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/registros-mantenimiento/ [get]
func (ctrl *GetAllRegistrosMantenimientoController) Run(c *gin.Context) {
	registros, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los registros de mantenimiento",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, registros)
}