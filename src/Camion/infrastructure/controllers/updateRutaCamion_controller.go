package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Camion/application"
	"github.com/vicpoo/API_recolecta/src/Camion/domain/entities"
)

type UpdateRutaCamionController struct {
	uc *application.UpdateRutaCamionUseCase
}

func NewUpdateRutaCamionController(
	uc *application.UpdateRutaCamionUseCase,
) *UpdateRutaCamionController {
	return &UpdateRutaCamionController{uc}
}

// @Summary      Actualizar ruta-camión
// @Tags         RutaCamion
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} map[string]interface{}
// @Router       /api/ruta-camion/{id} [put]
func (c *UpdateRutaCamionController) Run(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var rutaCamion entities.RutaCamion
	if err := ctx.ShouldBindJSON(&rutaCamion); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.uc.Execute(int32(id), &rutaCamion)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
