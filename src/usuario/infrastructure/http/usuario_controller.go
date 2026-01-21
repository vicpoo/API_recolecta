package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/usuario/application"
	"github.com/vicpoo/API_recolecta/src/usuario/domain"
	"github.com/vicpoo/API_recolecta/src/core"

)

type UsuarioController struct {
	create *application.CreateUsuario
	get    *application.GetUsuario
	list   *application.ListUsuarios
	login *application.LoginUsuario
	delete *application.DeleteUsuario
}

func NewUsuarioController(
	create *application.CreateUsuario,
	get *application.GetUsuario,
	list *application.ListUsuarios,
	login *application.LoginUsuario,
	delete *application.DeleteUsuario,
) *UsuarioController {
	return &UsuarioController{create, get, list, login, delete}
}

func (c *UsuarioController) RegisterRoutes(r *gin.Engine) {

	// ---------- RUTAS PÚBLICAS ----------
	public := r.Group("/usuarios")
	{
		public.POST("/login", c.Login)
		public.POST("", c.Create)
	}

	// ---------- RUTAS PROTEGIDAS ----------
	protected := r.Group(
		"/usuarios",
		core.JWTAuthMiddleware(),
	)

	{
		protected.GET(
			"",
			core.RequireRole(core.ADMIN, core.SUPERVISOR),
			c.List,
		)

		protected.GET(
			"/:id",
			core.RequireRole(core.ADMIN, core.SUPERVISOR),
			c.GetByID,
		)

		protected.DELETE(
			"/:id",
			core.RequireRole(core.ADMIN),
			c.Delete,
		)
	}
}


func (c *UsuarioController) Login(ctx *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.login.Execute(body.Email, body.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "credenciales inválidas"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (c *UsuarioController) Create(ctx *gin.Context) {
	var body domain.Usuario
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


func (c *UsuarioController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	usuario, err := c.get.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "usuario no encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, usuario)
}

func (c *UsuarioController) List(ctx *gin.Context) {
	usuarios, err := c.list.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, usuarios)
}

func (c *UsuarioController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))			

	if err := c.delete.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

