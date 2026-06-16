package controller_ciudadano

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Ciudadanos/application/application_ciudadano"
	"github.com/vicpoo/API_recolecta/src/core"
)

type CreateCiudadanoController struct {
	useCase *application_ciudadano.CreateCiudadano
}

func NewCreateCiudadanoController(useCase *application_ciudadano.CreateCiudadano) *CreateCiudadanoController {
	return &CreateCiudadanoController{useCase: useCase}
}

func (c *CreateCiudadanoController) Run(ctx *gin.Context) {
	var input application_ciudadano.CreateCiudadanoInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		core.RespondBadRequest(ctx, "json inválido", map[string]string{"detail": err.Error()})
		return
	}

	id, err := c.useCase.Execute(ctx.Request.Context(), input)
	if err != nil {
		core.RespondBadRequest(ctx, err.Error(), nil)
		return
	}

	core.RespondCreated(ctx, gin.H{"message": "ciudadano creado correctamente", "id": id})
}
