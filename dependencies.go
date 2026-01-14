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

	createUsuario := usuarioApp.NewCreateUsuario(usuarioRepo)
	getUsuario := usuarioApp.NewGetUsuario(usuarioRepo)
	listUsuarios := usuarioApp.NewListUsuarios(usuarioRepo)

	usuarioController := usuarioHTTP.NewUsuarioController(
		createUsuario,
		getUsuario,
		listUsuarios,
	)
	usuarioController.RegisterRoutes(engine)

	// ======================
	// START
	// ======================
	return engine.Run(":8080")
}
