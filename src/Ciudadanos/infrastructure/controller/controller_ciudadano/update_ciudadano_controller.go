package controller_ciudadano

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Ciudadanos/application/application_ciudadano"
	"github.com/vicpoo/API_recolecta/src/core"
)

type UpdateCiudadanoController struct {
	useCase *application_ciudadano.UpdateCiudadano
}

func NewUpdateCiudadanoController(useCase *application_ciudadano.UpdateCiudadano) *UpdateCiudadanoController {
	return &UpdateCiudadanoController{useCase: useCase}
}

func (c *UpdateCiudadanoController) Run(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		core.RespondInvalidInput(ctx, "id inválido")
		return
	}

	var input application_ciudadano.UpdateCiudadanoInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		core.RespondBadRequest(ctx, "json inválido", map[string]string{"detail": err.Error()})
		return
	}

	input.ID = id

	if err := c.useCase.Execute(ctx.Request.Context(), input); err != nil {
		core.RespondBadRequest(ctx, err.Error(), nil)
		return
	}

	core.RespondOK(ctx, gin.H{"message": "ciudadano actualizado correctamente"})
}
 