package controller_ciudadano

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Ciudadanos/application/application_ciudadano"
	"github.com/vicpoo/API_recolecta/src/core"
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
		core.RespondBadRequest(ctx, "json inválido", map[string]string{"detail": err.Error()})
		return
	}

	result, err := c.useCase.Execute(ctx.Request.Context(), input)
	if err != nil {
		core.RespondError(ctx, http.StatusUnauthorized, core.ErrCodeUnauthorized, "credenciales inválidas", nil)
		return
	}

	core.RespondOK(ctx, gin.H{"message": "login correcto", "token": result.Token, "data": result.Ciudadano})
}
