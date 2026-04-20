// GetAllSeguimientosFallaCriticaController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
)

type GetAllSeguimientosFallaCriticaController struct {
	getAllUseCase *application.GetAllSeguimientosFallaCriticaUseCase
}

func NewGetAllSeguimientosFallaCriticaController(getAllUseCase *application.GetAllSeguimientosFallaCriticaUseCase) *GetAllSeguimientosFallaCriticaController {
	return &GetAllSeguimientosFallaCriticaController{
		getAllUseCase: getAllUseCase,
	}
}

// @Summary      Listar seguimientos falla crítica
// @Tags         SeguimientoFallaCritica
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/seguimientos-falla-critica/ [get]
func (ctrl *GetAllSeguimientosFallaCriticaController) Run(c *gin.Context) {
	seguimientos, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener los seguimientos de falla crítica",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, seguimientos)
}