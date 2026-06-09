// GetIncidenciaByIDController.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/core"
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
		core.RespondInvalidInput(c, "ID de incidencia inválido")
		return
	}

	incidencia, err := ctrl.getByIDUseCase.Run(int32(id))
	if err != nil {
		core.RespondError(c, http.StatusNotFound, core.ErrCodeNotFound, "Incidencia no encontrada", nil)
		return
	}

	c.JSON(http.StatusOK, incidencia)
}
