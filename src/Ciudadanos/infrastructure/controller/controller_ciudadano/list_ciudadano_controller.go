package controller_ciudadano

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Ciudadanos/application/application_ciudadano"
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": ciudadanos,
	})
}