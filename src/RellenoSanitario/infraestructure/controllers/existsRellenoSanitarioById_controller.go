package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/RellenoSanitario/application"
)

type ExistsRellenoSanitarioByIdController struct {
	uc *application.ExistsRellenoSanitarioByIdUseCase
}

func NewExistsRellenoSanitarioByIdController(
	uc *application.ExistsRellenoSanitarioByIdUseCase,
) *ExistsRellenoSanitarioByIdController {
	return &ExistsRellenoSanitarioByIdController{uc: uc}
}

func (c *ExistsRellenoSanitarioByIdController) Execute(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	exists, err := c.uc.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":     id,
		"exists": exists,
	})
}
