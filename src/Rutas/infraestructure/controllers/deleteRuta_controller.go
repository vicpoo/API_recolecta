package controllers

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Rutas/application"
	"github.com/vicpoo/API_recolecta/src/core"
)

type DeleteRutaController struct {
	uc *application.DeleteRutaUseCase
}

func NewDeleteRutaController(uc *application.DeleteRutaUseCase) *DeleteRutaController {
	return &DeleteRutaController{uc}
}

// @Summary      Eliminar ruta
// @Tags         Ruta
// @Produce      json
// @Param        id path int true "ID ruta"
// @Success      200 {object} map[string]interface{}
// @Router       /api/rutas/{id} [delete]
func (ctr *DeleteRutaController) Run(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := ctr.uc.Run(int32(id))
	if err != nil {
		if strings.Contains(err.Error(), "no encontrado") || strings.Contains(err.Error(), "no encontrada") {
			core.RespondNotFound(ctx, "Ruta", ctx.Param("id"))
			return
		}
		core.RespondInternalServerError(ctx, "Error eliminando ruta", err)
		return
	}

	core.RespondOK(ctx, gin.H{"success": true, "message": "ruta eliminada"})
}
