package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	result, err := c.useCase.Execute(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "empleado no encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": result})
}
