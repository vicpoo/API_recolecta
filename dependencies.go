package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/vicpoo/API_recolecta/src/core"

	// colonia
	coloniaApp "github.com/vicpoo/API_recolecta/src/colonia/application"
	coloniaHTTP "github.com/vicpoo/API_recolecta/src/colonia/infrastructure/http"
	coloniaPG "github.com/vicpoo/API_recolecta/src/colonia/infrastructure/postgres"

	// usuario
	usuarioApp "github.com/vicpoo/API_recolecta/src/usuario/application"
	usuarioHTTP "github.com/vicpoo/API_recolecta/src/usuario/infrastructure/http"
	usuarioPG "github.com/vicpoo/API_recolecta/src/usuario/infrastructure/postgres"
)

func InitDependencies() error {
	// 1. ENV
	if err := godotenv.Load(); err != nil {
		return err
	}

	// 2. DB 
	db, err := core.ConnectPostgres()
	if err != nil {
		return err
	}

	// 3. Engine
	engine := gin.Default()
	engine.Use(core.CORSMiddleware())
	auth := engine.Group("/")
    auth.Use(core.JWTAuthMiddleware())
	
	// ======================
	// COLONIA
	// ======================
	coloniaRepo := coloniaPG.NewColoniaRepository(db)

	createColonia := coloniaApp.NewCreateColonia(coloniaRepo)
	getColonia := coloniaApp.NewGetColonia(coloniaRepo)
	listColonias := coloniaApp.NewListColonias(coloniaRepo)
	updateColonia := coloniaApp.NewUpdateColonia(coloniaRepo)

	coloniaController := coloniaHTTP.NewColoniaController(
		createColonia,
		getColonia,
		listColonias,
		updateColonia,
	)
	coloniaController.RegisterRoutes(engine)

	// ======================
	// USUARIO
	// ======================
	
	
	usuarioRepo := usuarioPG.NewUsuarioRepository(db)
	loginUsuario := usuarioApp.NewLoginUsuario(usuarioRepo)
	createUsuario := usuarioApp.NewCreateUsuario(usuarioRepo)

	getUsuario := usuarioApp.NewGetUsuario(usuarioRepo)
	listUsuarios := usuarioApp.NewListUsuarios(usuarioRepo)
	deleteUsuario := usuarioApp.NewDeleteUsuario(usuarioRepo)

	usuarioController := usuarioHTTP.NewUsuarioController(
		createUsuario,
		getUsuario,
		listUsuarios,
		loginUsuario,
		deleteUsuario,
	)
	usuarioController.RegisterRoutes(engine)


alertaRepo := alertaPG.NewAlertaRepository(db)

createAlerta := alertaApp.NewCreateAlerta(alertaRepo)
listAlertas := alertaApp.NewListMisAlertas(alertaRepo)
marcarLeida := alertaApp.NewMarcarLeida(alertaRepo)

alertaController := alertaHTTP.NewAlertaController(
	createAlerta,
	listAlertas,
	marcarLeida,
)

	
//=================
//Rutas protegidas
//=================

auth := engine.Group("/")
auth.Use(core.JWTAuthMiddleware())
alertaController.RegisterRoutes(auth)

admin := auth.Group("/colonias")
admin.Use(core.RequireRole(ADMIN, SUPERVISOR))

admin.POST("", coloniaController.Create)
admin.PUT("/:id", coloniaController.Update)
admin.DELETE("/:id", coloniaController.Delete)

auth.GET("/usuarios", usuarioController.List)
auth.GET("/usuarios/:id", usuarioController.GetByID)
auth.DELETE("/usuarios/:id", usuarioController.Delete)


auth.POST("/domicilios", domicilioController.Create)
auth.GET("/domicilios/:id", domicilioController.GetByID)
auth.PUT("/domicilios/:id", domicilioController.Update)
auth.DELETE("/domicilios/:id", domicilioController.Delete)

//============
//Publicas
//=========
	engine.POST("/usuarios", usuarioController.Create)
	engine.POST("/usuarios/login", usuarioController.Login)

	
	return engine.Run(":8080")
}
