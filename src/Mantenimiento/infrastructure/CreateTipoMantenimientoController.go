// CreateTipoMantenimientoController.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/application"
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/domain/entities"
)

type CreateTipoMantenimientoController struct {
	createUseCase *application.CreateTipoMantenimientoUseCase
}

func NewCreateTipoMantenimientoController(createUseCase *application.CreateTipoMantenimientoUseCase) *CreateTipoMantenimientoController {
	return &CreateTipoMantenimientoController{
		createUseCase: createUseCase,
	}
}

// @Summary      Crear tipo de mantenimiento
// @Tags         TipoMantenimiento
// @Produce      json
// @Param        body body entities.CreateRegistroMantenimientoRequest true "Body"
// @Success      200 {object} entities.RegistroMantenimientoResponse
// @Failure      400 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/tipos-mantenimiento/ [post]
func (ctrl *CreateTipoMantenimientoController) Run(c *gin.Context) {
	var request struct {
		Nombre    string `json:"nombre" binding:"required"`
		Categoria string `json:"categoria" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	// Crear el tipo de mantenimiento usando el constructor corregido
	tipoMantenimiento := entities.NewTipoMantenimiento(
		request.Nombre,
		request.Categoria,
	)

	createdTipoMantenimiento, err := ctrl.createUseCase.Run(tipoMantenimiento)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear el tipo de mantenimiento",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdTipoMantenimiento)
}
