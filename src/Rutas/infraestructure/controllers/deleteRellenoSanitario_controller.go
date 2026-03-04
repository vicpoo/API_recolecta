package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Rutas/application"
)

type DeleteRellenoSanitarioController struct {
	uc *application.DeleteRellenoSanitarioUseCase
}

func NewDeleteRellenoSanitarioController(uc *application.DeleteRellenoSanitarioUseCase) *DeleteRellenoSanitarioController {
	return &DeleteRellenoSanitarioController{uc: uc}
}

// @Summary      Eliminar relleno sanitario
// @Tags         RellenoSanitario
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} map[string]interface{}
// @Router       /api/relleno-sanitario/{id} [delete]
func (c *DeleteRellenoSanitarioController) Execute(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.uc.Execute(int32(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Relleno sanitario eliminado"})
}
