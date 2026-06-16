package controller_ciudadano

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Ciudadanos/application/application_ciudadano"
	"github.com/vicpoo/API_recolecta/src/core"
)

type ListCiudadanoController struct {
	useCase *application_ciudadano.ViewAllCiudadano
}

func NewListCiudadanoController(useCase *application_ciudadano.ViewAllCiudadano) *ListCiudadanoController {
	return &ListCiudadanoController{useCase: useCase}
}

func (c *ListCiudadanoController) Run(ctx *gin.Context) {
	ciudadanos, err := c.useCase.Execute(ctx.Request.Context())
	if err != nil {
		core.RespondInternalServerError(ctx, "no se pudo listar ciudadanos", err)
		return
	}

	core.RespondOK(ctx, gin.H{"data": ciudadanos})
}
