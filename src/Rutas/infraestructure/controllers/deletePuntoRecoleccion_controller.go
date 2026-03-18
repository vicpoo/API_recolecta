package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Rutas/application"
)

type DeletePuntoRecoleccionController struct {
	uc *application.DeletePuntoRecoleccionUseCase
}

func NewDeletePuntoRecoleccionController(uc *application.DeletePuntoRecoleccionUseCase) *DeletePuntoRecoleccionController {
	return &DeletePuntoRecoleccionController{uc: uc}
}

// @Summary      Eliminar punto de recolección
// @Tags         PuntoRecoleccion
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} map[string]interface{}
// @Router       /api/puntos-recoleccion/{id} [delete]
func (c *DeletePuntoRecoleccionController) Run(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	if err := c.uc.Execute(int32(id)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "punto de recolección eliminado"})
}
