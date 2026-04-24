package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/alerta_usuario/application"
	"github.com/vicpoo/API_recolecta/src/alerta_usuario/domain"
	"github.com/vicpoo/API_recolecta/src/core"
)

type AlertaController struct {
	create *application.CreateAlerta
	list   *application.ListMisAlertas
	read   *application.MarcarLeida
}

func NewAlertaController(
	create *application.CreateAlerta,
	list *application.ListMisAlertas,
	read *application.MarcarLeida,
) *AlertaController {
	return &AlertaController{create, list, read}
}

func (c *AlertaController) RegisterRoutes(r *gin.RouterGroup) {
	r.POST(
		"/alertas",
		core.RequireRole(core.SUPERVISOR),
		c.Create,
	)
	r.GET("/alertas", c.ListMine)
	r.PUT("/alertas/:id/leida", c.MarkAsRead)
}

// @Summary      Crear alerta de usuario
// @Tags         AlertaUsuario
// @Accept       json
// @Produce      json
// @Param        body body domain.AlertaUsuario true "Datos de la alerta"
// @Success      201
// @Failure      400 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Router       /api/alertas [post]
func (c *AlertaController) Create(ctx *gin.Context) {
	var body struct {
		Titulo    string `json:"titulo" binding:"required"`
		Mensaje   string `json:"mensaje" binding:"required"`
		UsuarioID int    `json:"usuario_id" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	alerta := domain.AlertaUsuario{
		Titulo:    body.Titulo,
		Mensaje:   body.Mensaje,
		UsuarioID: body.UsuarioID,           
		CreadoPor: ctx.GetInt("user_id"),    
	}

	if err := c.create.Execute(&alerta); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}

// @Summary      Listar mis alertas
// @Tags         AlertaUsuario
// @Produce      json
// @Success      200 {array} domain.AlertaUsuario
// @Failure      500 {object} map[string]interface{}
// @Router       /api/alertas [get]
func (c *AlertaController) ListMine(ctx *gin.Context) {
	alertas, err := c.list.Execute(ctx.GetInt("user_id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, alertas)
}

// @Summary      Marcar alerta como leída
// @Tags         AlertaUsuario
// @Produce      json
// @Param        id path int true "ID de la alerta"
// @Success      200
// @Failure      403 {object} map[string]interface{}
// @Router       /api/alertas/{id}/leida [put]
func (c *AlertaController) MarkAsRead(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	if err := c.read.Execute(id, ctx.GetInt("user_id")); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "no autorizado"})
		return
	}

	ctx.Status(http.StatusOK)
}