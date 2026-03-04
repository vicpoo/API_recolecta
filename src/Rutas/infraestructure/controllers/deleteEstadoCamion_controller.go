package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Rutas/application"
)

type DeleteEstadoCamionController struct {
	uc *application.DeleteEstadoCamionUseCase
}

func NewDeleteEstadoCamionController(
	uc *application.DeleteEstadoCamionUseCase,
) *DeleteEstadoCamionController {
	return &DeleteEstadoCamionController{
		uc: uc,
	}
}

// @Summary      Eliminar estado de camión
// @Tags         EstadoCamion
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} map[string]interface{}
// @Router       /api/estado-camion/{id} [delete]
func (ctr *DeleteEstadoCamionController) Run(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "id inválido",
		})
		return
	}

	if err := ctr.uc.Run(int32(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
