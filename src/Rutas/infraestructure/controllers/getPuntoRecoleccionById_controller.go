package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Rutas/application"
)

type GetPuntoRecoleccionByIdController struct {
	uc *application.GetPuntoRecoleccionByIdUseCase
}

func NewGetPuntoRecoleccionByIdController(uc *application.GetPuntoRecoleccionByIdUseCase) *GetPuntoRecoleccionByIdController {
	return &GetPuntoRecoleccionByIdController{uc: uc}
}

// @Summary      Punto de recolección por ID
// @Tags         PuntoRecoleccion
// @Produce      json
// @Param        id path int true "ID"
// @Success      200 {object} map[string]interface{}
// @Router       /api/puntos-recoleccion/{id} [get]
func (c *GetPuntoRecoleccionByIdController) Run(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	result, err := c.uc.Execute(int32(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
