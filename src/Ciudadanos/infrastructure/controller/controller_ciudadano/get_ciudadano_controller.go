package controller_ciudadano

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Ciudadanos/application/application_ciudadano"
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id inválido",
		})
		return
	}

	ciudadano, err := c.useCase.Execute(ctx.Request.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "ciudadano no encontrado",
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": ciudadano,
	})
}