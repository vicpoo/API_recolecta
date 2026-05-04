package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "json inválido", "detail": err.Error()})
		return
	}

	result, err := c.useCase.Execute(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "empleado creado correctamente", "data": result})
}
