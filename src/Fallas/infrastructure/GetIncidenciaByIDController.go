//GetIncidenciaByIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
)

type GetIncidenciaByIDController struct {
	getByIDUseCase *application.GetIncidenciaByIDUseCase
}

func NewGetIncidenciaByIDController(getByIDUseCase *application.GetIncidenciaByIDUseCase) *GetIncidenciaByIDController {
	return &GetIncidenciaByIDController{
		getByIDUseCase: getByIDUseCase,
	}
}

// @Summary      Incidencia por ID
// @Tags         Incidencia
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/incidencias/{id} [get]
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