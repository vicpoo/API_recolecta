package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Rutas/application"
)

type ExistsRellenoSanitarioByIdController struct {
	uc *application.ExistsRellenoSanitarioByIdUseCase
}

func NewExistsRellenoSanitarioByIdController(
	uc *application.ExistsRellenoSanitarioByIdUseCase,
) *ExistsRellenoSanitarioByIdController {
	return &ExistsRellenoSanitarioByIdController{uc: uc}
}

// @Summary      Verificar existencia relleno sanitario
// @Tags         RellenoSanitario
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} map[string]interface{}
// @Router       /api/relleno-sanitario/exists/{id} [get]
func (c *ExistsRellenoSanitarioByIdController) Execute(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	exists, err := c.uc.Execute(int32(id))
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
