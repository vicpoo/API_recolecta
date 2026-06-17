// CreateSeguimientoFallaCriticaController.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/Fallas/domain/entities"
	"github.com/vicpoo/API_recolecta/src/core"
)

type CreateSeguimientoFallaCriticaController struct {
	createUseCase *application.CreateSeguimientoFallaCriticaUseCase
}

func NewCreateSeguimientoFallaCriticaController(createUseCase *application.CreateSeguimientoFallaCriticaUseCase) *CreateSeguimientoFallaCriticaController {
	return &CreateSeguimientoFallaCriticaController{
		createUseCase: createUseCase,
	}
}

// @Summary      Crear seguimiento falla crítica
// @Tags         SeguimientoFallaCritica
// @Produce      json
// @Param        body body entities.CreateSeguimientoFallaCriticaRequest true "Body"
// @Success      200 {object} entities.SeguimientoFallaCriticaResponse
// @Failure      400 {object} core.ErrorResponse
// @Router       /api/seguimientos-falla-critica/ [post]
func (ctrl *CreateSeguimientoFallaCriticaController) Run(c *gin.Context) {
	var request struct {
		FallaID    int32  `json:"falla_id" binding:"required"`
		Comentario string `json:"comentario" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		core.RespondBadRequest(c, "Datos inválidos", map[string]string{"error": err.Error()})
		return
	}

	seguimiento := entities.NewSeguimientoFallaCritica(request.FallaID, request.Comentario)

	createdSeguimiento, err := ctrl.createUseCase.Run(seguimiento)
	if err != nil {
		core.RespondInternalServerError(c, "No se pudo crear el seguimiento de falla crítica", err)
		return
	}

	core.RespondCreated(c, createdSeguimiento)
}
