package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Rutas/application"
)

type GetAllRellenoSanitarioController struct {
	uc *application.ListRellenoSanitarioUseCase
}

func NewGetAllRellenoSanitarioController(uc *application.ListRellenoSanitarioUseCase) *GetAllRellenoSanitarioController {
	return &GetAllRellenoSanitarioController{uc: uc}
}

// @Summary      Listar rellenos sanitarios
// @Tags         RellenoSanitario
// @Produce      json
// @Success      200 {array} map[string]interface{}
// @Router       /api/relleno-sanitario/ [get]
func (c *GetAllRellenoSanitarioController) Execute(ctx *gin.Context) {
	result, err := c.uc.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
