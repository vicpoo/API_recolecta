//GetAllIncidenciasController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
)

type GetAllIncidenciasController struct {
	getAllUseCase *application.GetAllIncidenciasUseCase
}

func NewGetAllIncidenciasController(getAllUseCase *application.GetAllIncidenciasUseCase) *GetAllIncidenciasController {
	return &GetAllIncidenciasController{
		getAllUseCase: getAllUseCase,
	}
}

// @Summary      Listar incidencias
// @Tags         Incidencia
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/incidencias/ [get]
func (ctrl *GetAllIncidenciasController) Run(c *gin.Context) {
	incidencias, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las incidencias",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, incidencias)
}