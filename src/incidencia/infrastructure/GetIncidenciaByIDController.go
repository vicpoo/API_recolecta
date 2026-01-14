//GetIncidenciaByIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/incidencia/application"
)

type GetIncidenciaByIDController struct {
	getByIDUseCase *application.GetIncidenciaByIDUseCase
}

func NewGetIncidenciaByIDController(getByIDUseCase *application.GetIncidenciaByIDUseCase) *GetIncidenciaByIDController {
	return &GetIncidenciaByIDController{
		getByIDUseCase: getByIDUseCase,
	}
}

func (ctrl *GetIncidenciaByIDController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	incidencia, err := ctrl.getByIDUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No se pudo encontrar la incidencia",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, incidencia)
}