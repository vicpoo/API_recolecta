package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Rutas/application"
)

type GetAllPuntoRecoleccionController struct {
	uc *application.ListAllPuntoRecoleccionUseCase
}

func NewGetAllPuntoRecoleccionController(uc *application.ListAllPuntoRecoleccionUseCase) *GetAllPuntoRecoleccionController {
	return &GetAllPuntoRecoleccionController{uc: uc}
}

// @Summary      Listar puntos de recolección
// @Tags         PuntoRecoleccion
// @Produce      json
// @Success      200 {array} map[string]interface{}
// @Router       /api/puntos-recoleccion/ [get]
func (c *GetAllPuntoRecoleccionController) Run(ctx *gin.Context) {
	result, err := c.uc.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
