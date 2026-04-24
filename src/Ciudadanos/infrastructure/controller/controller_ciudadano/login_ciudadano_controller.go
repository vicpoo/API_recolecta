package controller_ciudadano

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Ciudadanos/application/application_ciudadano"
)

type LoginCiudadanoController struct {
	useCase *application_ciudadano.LoginCiudadano
}

func NewLoginCiudadanoController(useCase *application_ciudadano.LoginCiudadano) *LoginCiudadanoController {
	return &LoginCiudadanoController{useCase: useCase}
}

func (c *LoginCiudadanoController) Run(ctx *gin.Context) {
	var input application_ciudadano.LoginCiudadanoInput

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
		"data":    result.Ciudadano,
	})
}