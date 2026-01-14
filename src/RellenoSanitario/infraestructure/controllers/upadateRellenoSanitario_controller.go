package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/RellenoSanitario/application"
	"github.com/vicpoo/API_recolecta/src/RellenoSanitario/domain/entities"
)

type UpdateRellenoSanitarioController struct {
	uc *application.UpdateRellenoSanitarioUseCase
}

func NewUpdateRellenoSanitarioController(uc *application.UpdateRellenoSanitarioUseCase) *UpdateRellenoSanitarioController {
	return &UpdateRellenoSanitarioController{uc: uc}
}

func (c *UpdateRellenoSanitarioController) Execute(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var relleno entities.RellenoSanitario
	if err := ctx.ShouldBindJSON(&relleno); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.uc.Execute(int32(id), &relleno)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
