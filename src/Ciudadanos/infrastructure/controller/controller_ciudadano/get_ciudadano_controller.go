package controller_ciudadano

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Ciudadanos/application/application_ciudadano"
	"github.com/vicpoo/API_recolecta/src/core"
)

type GetCiudadanoController struct {
	useCase *application_ciudadano.ViewOneCiudadano
}

func NewGetCiudadanoController(useCase *application_ciudadano.ViewOneCiudadano) *GetCiudadanoController {
	return &GetCiudadanoController{useCase: useCase}
}

func (c *GetCiudadanoController) Run(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondInvalidInput(ctx, "id inválido")
		return
	}

	ciudadano, err := c.useCase.Execute(ctx.Request.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows || err.Error() == "ciudadano no encontrado" {
			core.RespondNotFound(ctx, "ciudadano", strconv.Itoa(id))
			return
		}

		core.RespondInternalServerError(ctx, "no se pudo obtener ciudadano", err)
		return
	}

	core.RespondOK(ctx, gin.H{"data": ciudadano})
}
