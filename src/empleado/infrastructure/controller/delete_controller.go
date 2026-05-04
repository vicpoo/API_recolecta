package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	if err := c.useCase.Execute(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "empleado eliminado correctamente"})
}
