package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	var input application.UpdateEmpleadoInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "json inválido", "detail": err.Error()})
		return
	}

	input.ID = id

	result, err := c.useCase.Execute(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "empleado actualizado correctamente", "data": result})
}
