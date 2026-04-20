package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Rutas/application"
)

type GetRellenoSanitarioByNombreController struct {
	uc *application.GetRellenoSanitarioByNombreUseCase
}

func NewGetRellenoSanitarioByNombreController(
	uc *application.GetRellenoSanitarioByNombreUseCase,
) *GetRellenoSanitarioByNombreController {
	return &GetRellenoSanitarioByNombreController{uc: uc}
}

// @Summary      Buscar relleno sanitario por nombre
// @Tags         RellenoSanitario
// @Produce      json
// @Param        nombre query string true "Nombre"
// @Success      200 {object} map[string]interface{}
// @Router       /api/relleno-sanitario/buscar [get]
func (c *GetRellenoSanitarioByNombreController) Execute(ctx *gin.Context) {
	nombre := ctx.Query("nombre")

	if nombre == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "El parámetro 'nombre' es requerido",
		})
		return
	}

	result, err := c.uc.Execute(nombre)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
