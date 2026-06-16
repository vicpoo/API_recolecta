package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/core"
	"github.com/vicpoo/API_recolecta/src/empleado/application"
)

type DeleteEmpleadoController struct {
	useCase *application.DeleteEmpleado
}

func NewDeleteEmpleadoController(useCase *application.DeleteEmpleado) *DeleteEmpleadoController {
	return &DeleteEmpleadoController{useCase: useCase}
}

func (c *DeleteEmpleadoController) Run(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		core.RespondBadRequest(ctx, "id inválido", nil)
		return
	}

	if err := c.useCase.Execute(ctx.Request.Context(), id); err != nil {
		if err.Error() == "empleado no encontrado" {
			core.RespondNotFound(ctx, "empleado", strconv.Itoa(id))
			return
		}
		core.RespondBadRequest(ctx, err.Error(), nil)
		return
	}

	core.RespondOK(ctx, gin.H{"message": "empleado eliminado correctamente"})
}
