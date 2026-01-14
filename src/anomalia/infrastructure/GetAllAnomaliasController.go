// GetAllAnomaliasController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/anomalia/application"
)

type GetAllAnomaliasController struct {
	getAllUseCase *application.GetAllAnomaliasUseCase
}

func NewGetAllAnomaliasController(getAllUseCase *application.GetAllAnomaliasUseCase) *GetAllAnomaliasController {
	return &GetAllAnomaliasController{
		getAllUseCase: getAllUseCase,
	}
}

func (ctrl *GetAllAnomaliasController) Run(c *gin.Context) {
	anomalias, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudieron obtener las anomalías",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, anomalias)
}