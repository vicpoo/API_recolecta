package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/empleado/application"
)

type ListEmpleadoController struct {
	useCase *application.ListEmpleado
}

func NewListEmpleadoController(useCase *application.ListEmpleado) *ListEmpleadoController {
	return &ListEmpleadoController{useCase: useCase}
}

func (c *ListEmpleadoController) Run(ctx *gin.Context) {
	result, err := c.useCase.Execute(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": result})
}
