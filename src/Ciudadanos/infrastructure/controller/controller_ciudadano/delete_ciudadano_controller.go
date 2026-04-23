package controller_ciudadano

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Ciudadanos/application/application_ciudadano"
)

type DeleteCiudadanoController struct {
	useCase *application_ciudadano.DeleteCiudadano
}

func NewDeleteCiudadanoController(useCase *application_ciudadano.DeleteCiudadano) *DeleteCiudadanoController {
	return &DeleteCiudadanoController{useCase: useCase}
}

func (c *DeleteCiudadanoController) Run(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id inválido",
		})
		return
	}

	err = c.useCase.Execute(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ciudadano eliminado correctamente",
	})
}