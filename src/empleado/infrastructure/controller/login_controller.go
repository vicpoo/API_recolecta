package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vicpoo/API_recolecta/src/empleado/application"
)

type LoginEmpleadoController struct {
	useCase *application.LoginEmpleado
}

func NewLoginEmpleadoController(useCase *application.LoginEmpleado) *LoginEmpleadoController {
	return &LoginEmpleadoController{useCase: useCase}
}

func (c *LoginEmpleadoController) Run(ctx *gin.Context) {
	var input application.LoginEmpleadoInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":  "json inválido",
			"detail": err.Error(),
		})
		return
	}

	result, err := c.useCase.Execute(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "login correcto",
		"token":   result.Token,
		"data":    result.Empleado,
	})
}
