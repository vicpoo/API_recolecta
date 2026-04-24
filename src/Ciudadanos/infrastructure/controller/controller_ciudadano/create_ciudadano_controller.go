package controller_ciudadano

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Ciudadanos/application/application_ciudadano"
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "json inválido",
			"detail": err.Error(),
		})
		return
	}

	id, err := c.useCase.Execute(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "ciudadano creado correctamente",
		"id":      id,
	})
}