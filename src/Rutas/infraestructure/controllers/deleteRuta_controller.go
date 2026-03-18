package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Rutas/application"
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
		ctx.JSON(http.StatusNotFound, gin.H{"success": false, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "message": "ruta eliminada"})
}
