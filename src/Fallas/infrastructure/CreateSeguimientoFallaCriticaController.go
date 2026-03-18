// CreateSeguimientoFallaCriticaController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Fallas/application"
	"github.com/vicpoo/API_recolecta/src/Fallas/domain/entities"
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
// @Param        body body map[string]interface{} true "Body"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /api/seguimientos-falla-critica/ [post]
func (ctrl *CreateSeguimientoFallaCriticaController) Run(c *gin.Context) {
	var request struct {
		FallaID    int32  `json:"falla_id" binding:"required"`
		Comentario string `json:"comentario" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	seguimiento := entities.NewSeguimientoFallaCritica(request.FallaID, request.Comentario)

	createdSeguimiento, err := ctrl.createUseCase.Run(seguimiento)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear el seguimiento de falla crítica",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdSeguimiento)
}