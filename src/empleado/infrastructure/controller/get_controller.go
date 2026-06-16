package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/core"
	"github.com/vicpoo/API_recolecta/src/empleado/application"
)

type GetEmpleadoController struct {
	useCase *application.GetEmpleado
}

func NewGetEmpleadoController(useCase *application.GetEmpleado) *GetEmpleadoController {
	return &GetEmpleadoController{useCase: useCase}
}

func (c *GetEmpleadoController) Run(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		core.RespondBadRequest(ctx, "id inválido", nil)
		return
	}

	result, err := c.useCase.Execute(ctx.Request.Context(), id)
	if err != nil {
		if err.Error() == "empleado no encontrado" {
			core.RespondNotFound(ctx, "empleado", strconv.Itoa(id))
			return
		}
		core.RespondInternalServerError(ctx, "error obteniendo empleado", err)
		return
	}

	if result == nil {
		core.RespondNotFound(ctx, "empleado", strconv.Itoa(id))
		return
	}

	core.RespondOK(ctx, gin.H{"data": result})
}
