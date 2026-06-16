// GetRegistroMantenimientoByIDController.go
package infrastructure

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetRegistroMantenimientoByIDController struct {
	getByIDUseCase *application.GetRegistroMantenimientoByIDUseCase
}

func NewGetRegistroMantenimientoByIDController(getByIDUseCase *application.GetRegistroMantenimientoByIDUseCase) *GetRegistroMantenimientoByIDController {
	return &GetRegistroMantenimientoByIDController{
		getByIDUseCase: getByIDUseCase,
	}
}

// @Summary      Registro por ID
// @Tags         RegistroMantenimiento
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/registros-mantenimiento/{id} [get]
func (ctrl *GetRegistroMantenimientoByIDController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondInvalidInput(c, "ID inválido")
		return
	}

	registro, err := ctrl.getByIDUseCase.Run(int32(id))
	if err != nil {
		if strings.Contains(err.Error(), "no encontrado") {
			core.RespondNotFound(c, "Registro de mantenimiento", idParam)
		} else {
			core.RespondInternalServerError(c, "No se pudo obtener el registro de mantenimiento", err)
		}
		return
	}

	core.RespondOK(c, registro)
}
