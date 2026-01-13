//GetAllIncidenciasController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/incidencia/application"
)

type GetAllIncidenciasController struct {
	getAllUseCase *application.GetAllIncidenciasUseCase
}

func NewGetAllIncidenciasController(getAllUseCase *application.GetAllIncidenciasUseCase) *GetAllIncidenciasController {
	return &GetAllIncidenciasController{
		getAllUseCase: getAllUseCase,
	}
}

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