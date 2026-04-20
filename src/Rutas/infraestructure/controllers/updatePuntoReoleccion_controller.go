package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Rutas/application"
	"github.com/vicpoo/API_recolecta/src/Rutas/domain/entities"
)

type UpdatePuntoRecoleccionController struct {
	uc *application.UpdatePuntoRecoleccionUseCase
}

func NewUpdatePuntoRecoleccionController(uc *application.UpdatePuntoRecoleccionUseCase) *UpdatePuntoRecoleccionController {
	return &UpdatePuntoRecoleccionController{uc: uc}
}

// @Summary      Actualizar punto de recolección
// @Tags         PuntoRecoleccion
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} map[string]interface{}
// @Router       /api/puntos-recoleccion/{id} [put]
func (c *UpdatePuntoRecoleccionController) Run(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	var p entities.PuntoRecoleccion
	if err := ctx.ShouldBindJSON(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.uc.Execute(int32(id), &p)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
