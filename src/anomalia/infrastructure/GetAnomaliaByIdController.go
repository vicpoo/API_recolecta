// GetAnomaliaByIdController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/anomalia/application"
)

type GetAnomaliaByIdController struct {
	getByIdUseCase *application.GetAnomaliaByIdUseCase
}

func NewGetAnomaliaByIdController(getByIdUseCase *application.GetAnomaliaByIdUseCase) *GetAnomaliaByIdController {
	return &GetAnomaliaByIdController{
		getByIdUseCase: getByIdUseCase,
	}
}

func (ctrl *GetAnomaliaByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	anomalia, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo obtener la anomalía",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, anomalia)
}