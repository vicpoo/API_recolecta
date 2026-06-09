package controller

import (

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/core"
	"github.com/vicpoo/API_recolecta/src/empleado/application"
)

type CreateEmpleadoController struct {
	useCase *application.CreateEmpleado
}

func NewCreateEmpleadoController(useCase *application.CreateEmpleado) *CreateEmpleadoController {
	return &CreateEmpleadoController{useCase: useCase}
}

func (c *CreateEmpleadoController) Run(ctx *gin.Context) {
	var input application.CreateEmpleadoInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		core.RespondBadRequest(ctx, "json inválido", err.Error())
		return
	}

	result, err := c.useCase.Execute(ctx.Request.Context(), input)
	if err != nil {
		core.RespondBadRequest(ctx, err.Error(), nil)
		return
	}

	core.RespondCreated(ctx, gin.H{"message": "empleado creado correctamente", "data": result})
}
