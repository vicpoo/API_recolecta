// GetAllIncidenciasController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/core"
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
// @Success      200 {object} entities.IncidenciaListResponse
// @Failure      400 {object} core.ErrorResponse
// @Router       /api/incidencias/ [get]
func (ctrl *GetAllIncidenciasController) Run(c *gin.Context) {
	incidencias, err := ctrl.getAllUseCase.Run()
	if err != nil {
		core.RespondInternalServerError(c, "Error al listar incidencias", err)
		return
	}

	c.JSON(http.StatusOK, incidencias)
}
