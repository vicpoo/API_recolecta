package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Rutas/application"
	"github.com/vicpoo/API_recolecta/src/Rutas/domain/entities"
	"github.com/vicpoo/API_recolecta/src/core"
)

type CreatePuntoRecoleccionController struct {
	uc *application.SavePuntoRecoleccionUseCase
}

func NewCreatePuntoRecoleccionController(uc *application.SavePuntoRecoleccionUseCase) *CreatePuntoRecoleccionController {
	return &CreatePuntoRecoleccionController{uc: uc}
}

// @Summary      Crear punto de recolección
// @Tags         PuntoRecoleccion
// @Accept       json
// @Produce      json
// @Success      201 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/puntos-recoleccion/ [post]
func (c *CreatePuntoRecoleccionController) Run(ctx *gin.Context) {
	var p entities.PuntoRecoleccion

	if err := ctx.ShouldBindJSON(&p); err != nil {
		core.RespondBadRequest(ctx, "body inválido", map[string]string{"error": err.Error()})
		return
	}

	result, err := c.uc.Execute(&p)
	if err != nil {
		core.RespondInternalServerError(ctx, "No se pudo crear el punto de recolección", err)
		return
	}

	core.RespondCreated(ctx, result)
}
