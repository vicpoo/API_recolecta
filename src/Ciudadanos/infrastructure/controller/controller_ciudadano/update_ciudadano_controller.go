package controller_ciudadano

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Ciudadanos/application/application_ciudadano"
)

type UpdateCiudadanoController struct {
	useCase *application_ciudadano.UpdateCiudadano
}

func NewUpdateCiudadanoController(useCase *application_ciudadano.UpdateCiudadano) *UpdateCiudadanoController {
	return &UpdateCiudadanoController{useCase: useCase}
}

func (c *UpdateCiudadanoController) Run(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id inválido",
		})
		return
	}

	var input application_ciudadano.UpdateCiudadanoInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "json inválido",
			"detail": err.Error(),
		})
		return
	}

	input.ID = id

	err = c.useCase.Execute(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ciudadano actualizado correctamente",
	})
}