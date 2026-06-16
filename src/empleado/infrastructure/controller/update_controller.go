package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/core"
	"github.com/vicpoo/API_recolecta/src/empleado/application"
)

type UpdateEmpleadoController struct {
	useCase *application.UpdateEmpleado
}

func NewUpdateEmpleadoController(useCase *application.UpdateEmpleado) *UpdateEmpleadoController {
	return &UpdateEmpleadoController{useCase: useCase}
}

func (c *UpdateEmpleadoController) Run(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		core.RespondBadRequest(ctx, "id inválido", nil)
		return
	}

	var input application.UpdateEmpleadoInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		core.RespondBadRequest(ctx, "json inválido", err.Error())
		return
	}

	input.ID = id

	result, err := c.useCase.Execute(ctx.Request.Context(), input)
	if err != nil {
		if err.Error() == "empleado no encontrado" {
			core.RespondNotFound(ctx, "empleado", strconv.Itoa(id))
			return
		}
		core.RespondBadRequest(ctx, err.Error(), nil)
		return
	}

	core.RespondOK(ctx, gin.H{"message": "empleado actualizado correctamente", "data": result})
}
