package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/rol/application"
)

type RoleController struct {
	create *application.CreateRole
	get    *application.GetRole
	list   *application.ListRoles
	update *application.UpdateRole
}

func NewRoleController(
	create *application.CreateRole,
	get *application.GetRole,
	list *application.ListRoles,
	update *application.UpdateRole,
) *RoleController {
	return &RoleController{create, get, list, update}
}

func (c *RoleController) RegisterRoutes(r *gin.Engine) {
	group := r.Group("/roles")
	{
		group.POST("", c.Create)
		group.GET("", c.List)
		group.GET("/:id", c.GetByID)
		group.PUT("/:id", c.Update)
	}
}

func (c *RoleController) Create(ctx *gin.Context) {
	var body struct {
		Nombre string `json:"nombre"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.create.Execute(body.Nombre)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusCreated)
}

func (c *RoleController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	role, err := c.get.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "rol no encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, role)
}

func (c *RoleController) List(ctx *gin.Context) {
	roles, err := c.list.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, roles)
}

func (c *RoleController) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var body struct {
		Nombre string `json:"nombre"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.update.Execute(id, body.Nombre)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}
