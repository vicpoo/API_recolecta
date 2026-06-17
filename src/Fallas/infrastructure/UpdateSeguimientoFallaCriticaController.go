// UpdateSeguimientoFallaCriticaController.go
package infrastructure

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/Fallas/domain/entities"
	"github.com/vicpoo/API_recolecta/src/core"
)

type UpdateSeguimientoFallaCriticaController struct {
	updateUseCase *application.UpdateSeguimientoFallaCriticaUseCase
}

func NewUpdateSeguimientoFallaCriticaController(updateUseCase *application.UpdateSeguimientoFallaCriticaUseCase) *UpdateSeguimientoFallaCriticaController {
	return &UpdateSeguimientoFallaCriticaController{
		updateUseCase: updateUseCase,
	}
}

// @Summary      Actualizar seguimiento falla crítica
// @Tags         SeguimientoFallaCritica
// @Produce      json
// @Param        body body entities.UpdateSeguimientoFallaCriticaRequest true "Body"
// @Param        id path int true "ID"
// @Success      200 {object} entities.SeguimientoFallaCriticaResponse
// @Failure      400 {object} core.ErrorResponse
// @Router       /api/seguimientos-falla-critica/{id} [put]
func (ctrl *UpdateSeguimientoFallaCriticaController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondBadRequest(c, "ID inválido", map[string]string{"error": err.Error()})
		return
	}

	var request struct {
		FallaID    int32  `json:"falla_id" binding:"required"`
		Comentario string `json:"comentario" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		core.RespondBadRequest(c, "Datos inválidos", map[string]string{"error": err.Error()})
		return
	}

	seguimiento := &entities.SeguimientoFallaCritica{
		SeguimientoID: int32(id),
		FallaID:       request.FallaID,
		Comentario:    request.Comentario,
	}

	updatedSeguimiento, err := ctrl.updateUseCase.Run(seguimiento)
	if err != nil {
		core.RespondInternalServerError(c, "No se pudo actualizar el seguimiento de falla crítica", err)
		return
	}

	core.RespondOK(c, updatedSeguimiento)
}
