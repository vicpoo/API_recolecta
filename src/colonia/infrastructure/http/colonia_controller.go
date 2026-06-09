package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/colonia/application"
	"github.com/vicpoo/API_recolecta/src/colonia/domain"
	"github.com/vicpoo/API_recolecta/src/core"
)

type ColoniaController struct {
	create *application.CreateColonia
	get    *application.GetColonia
	list   *application.ListColonias
	update *application.UpdateColonia
	delete *application.DeleteColonia
}

func NewColoniaController(
	create *application.CreateColonia,
	get *application.GetColonia,
	list *application.ListColonias,
	update *application.UpdateColonia,
	delete *application.DeleteColonia,
) *ColoniaController {
	return &ColoniaController{create, get, list, update, delete}
}

func (c *ColoniaController) RegisterRoutes(r *gin.Engine) {

	public := r.Group("/api/colonia")
	{
		public.GET("", c.List)
		public.GET("/:id", c.GetByID)
	}

	// Rutas protegidas SOLO ADMIN
	admin := r.Group(
		"/api/colonia",
		core.JWTAuthMiddleware(),
		core.RequireRole(core.ADMIN),
	)
	{
		admin.POST("", c.Create)
		admin.PUT("/:id", c.Update)
		admin.DELETE("/:id", c.Delete)
	}
}

// @Summary      Crear colonia
// @Description  Crea una nueva colonia. Solo administradores (rol ADMIN)
// @Tags         Colonia
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        body body domain.Colonia true "Datos de la colonia"
// @Success      201 {object} domain.Colonia
// @Failure      400 {object} map[string]interface{} "Datos inválidos"
// @Failure      401 {object} map[string]interface{} "No autenticado"
// @Failure      403 {object} map[string]interface{} "No autorizado (requiere rol ADMIN)"
// @Failure      500 {object} map[string]interface{} "Error interno del servidor"
// @Router       /api/colonia [post]
func (c *ColoniaController) Create(ctx *gin.Context) {
	var body domain.Colonia
	if err := ctx.ShouldBindJSON(&body); err != nil {
		core.RespondValidationError(ctx, "Datos de colonia inválidos", map[string]string{"error": err.Error()})
		return
	}

	if err := c.create.Execute(&body); err != nil {
		core.RespondInternalServerError(ctx, "Error al crear colonia", err)
		return
	}

	ctx.JSON(http.StatusCreated, body)
}

// @Summary      Obtener colonia por ID
// @Description  Obtiene los detalles de una colonia específica. Endpoint público
// @Tags         Colonia
// @Produce      json
// @Param        id path int true "ID de la colonia"
// @Success      200 {object} domain.Colonia
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      404 {object} map[string]interface{} "Colonia no encontrada"
// @Failure      500 {object} map[string]interface{} "Error interno del servidor"
// @Router       /api/colonia/{id} [get]
func (c *ColoniaController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		core.RespondInvalidInput(ctx, "ID de colonia inválido")
		return
	}

	colonia, err := c.get.Execute(id)
	if err != nil {
		core.RespondError(ctx, http.StatusNotFound, core.ErrCodeNotFound, "Colonia no encontrada", nil)
		return
	}

	ctx.JSON(http.StatusOK, colonia)
}

// @Summary      Listar colonias
// @Description  Obtiene el listado de todas las colonias disponibles. Endpoint público
// @Tags         Colonia
// @Produce      json
// @Success      200 {array} domain.Colonia
// @Failure      500 {object} map[string]interface{} "Error interno del servidor"
// @Router       /api/colonia [get]
func (c *ColoniaController) List(ctx *gin.Context) {
	colonias, err := c.list.Execute()
	if err != nil {
		core.RespondInternalServerError(ctx, "Error al listar colonias", err)
		return
	}

	ctx.JSON(http.StatusOK, colonias)
}

// @Summary      Actualizar colonia
// @Description  Actualiza los datos de una colonia existente. Solo administradores (rol ADMIN)
// @Tags         Colonia
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        id path int true "ID de la colonia"
// @Param        body body domain.Colonia true "Datos a actualizar"
// @Success      200 "Colonia actualizada correctamente"
// @Failure      400 {object} map[string]interface{} "Datos inválidos"
// @Failure      401 {object} map[string]interface{} "No autenticado"
// @Failure      403 {object} map[string]interface{} "No autorizado (requiere rol ADMIN)"
// @Failure      500 {object} map[string]interface{} "Error interno del servidor"
// @Router       /api/colonia/{id} [put]
func (c *ColoniaController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		core.RespondInvalidInput(ctx, "ID de colonia inválido")
		return
	}

	var body domain.Colonia
	if err := ctx.ShouldBindJSON(&body); err != nil {
		core.RespondValidationError(ctx, "Datos de colonia inválidos", map[string]string{"error": err.Error()})
		return
	}

	body.ColoniaID = id

	if err := c.update.Execute(&body); err != nil {
		core.RespondInternalServerError(ctx, "Error al actualizar colonia", err)
		return
	}

	ctx.Status(http.StatusOK)
}

// @Summary      Eliminar colonia
// @Description  Elimina una colonia de la base de datos. Solo administradores (rol ADMIN)
// @Tags         Colonia
// @Produce      json
// @Security     Bearer
// @Param        id path int true "ID de la colonia"
// @Success      204 "Colonia eliminada correctamente"
// @Failure      400 {object} map[string]interface{} "ID inválido"
// @Failure      401 {object} map[string]interface{} "No autenticado"
// @Failure      403 {object} map[string]interface{} "No autorizado (requiere rol ADMIN)"
// @Failure      500 {object} map[string]interface{} "Error interno del servidor"
// @Router       /api/colonia/{id} [delete]
func (c *ColoniaController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		core.RespondInvalidInput(ctx, "ID de colonia inválido")
		return
	}

	if err := c.delete.Execute(id); err != nil {
		core.RespondInternalServerError(ctx, "Error al eliminar colonia", err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
