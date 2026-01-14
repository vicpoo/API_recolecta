package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
	"github.com/vicpoo/API_recolecta/src/notificacion/domain"
)

type NotificacionController struct {
	create     *application.CreateNotificacion
	get        *application.GetNotificacion
	listUser   *application.ListNotificacionesUsuario
	deactivate *application.DeactivateNotificacion
}

func NewNotificacionController(
	create *application.CreateNotificacion,
	get *application.GetNotificacion,
	listUser *application.ListNotificacionesUsuario,
	deactivate *application.DeactivateNotificacion,
) *NotificacionController {
	return &NotificacionController{create, get, listUser, deactivate}
}

func (c *NotificacionController) RegisterRoutes(r *gin.Engine) {
	group := r.Group("/notificaciones")
	{
		group.POST("", c.Create)
		group.GET("/:id", c.GetByID)
		group.GET("/usuario/:usuarioId", c.ListByUsuario)
		group.PUT("/:id/desactivar", c.Deactivate)
	}
}

func (c *NotificacionController) Create(ctx *gin.Context) {
	var body domain.Notificacion
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.create.Execute(&body); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (c *NotificacionController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	notificacion, err := c.get.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "notificación no encontrada"})
		return
	}
	ctx.JSON(http.StatusOK, notificacion)
}

func (c *NotificacionController) ListByUsuario(ctx *gin.Context) {
	usuarioID, _ := strconv.Atoi(ctx.Param("usuarioId"))
	notificaciones, err := c.listUser.Execute(usuarioID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, notificaciones)
}

func (c *NotificacionController) Deactivate(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.deactivate.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}
