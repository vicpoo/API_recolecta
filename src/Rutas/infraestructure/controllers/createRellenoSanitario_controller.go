package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Rutas/application"
	"github.com/vicpoo/API_recolecta/src/Rutas/domain/entities"
)

type CreateRellenoSanitarioController struct {
	uc *application.SaveRellenoSanitarioUseCase
}

func NewCreateRellenoSanitarioController(uc *application.SaveRellenoSanitarioUseCase) *CreateRellenoSanitarioController {
	return &CreateRellenoSanitarioController{uc: uc}
}

func (c *CreateRellenoSanitarioController) Execute(ctx *gin.Context) {
	var relleno entities.RellenoSanitario

	if err := ctx.ShouldBindJSON(&relleno); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.uc.Execute(&relleno)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, result)
}
