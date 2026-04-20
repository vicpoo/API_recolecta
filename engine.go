package main

import (
	"time"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/vicpoo/API_recolecta/docs"
	"github.com/vicpoo/API_recolecta/src/core"

	"github.com/gin-gonic/gin"

	historialRoutes "github.com/vicpoo/API_recolecta/src/Camion/infrastructure/routes"
	rutaCamionRoutes "github.com/vicpoo/API_recolecta/src/Camion/infrastructure/routes"
	tipoCamionRoutes "github.com/vicpoo/API_recolecta/src/Camion/infrastructure/routes"

	camionRoutes "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/routes"
	estadoCamionRoutes "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/routes"
	puntoRoutes "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/routes"
	registroVaciadoRoutes "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/routes"
	rsRoutes "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/routes"
	rutaRoutes "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/routes"

	alertaMantenimientoInfra "github.com/vicpoo/API_recolecta/src/Mantenimiento/infrastructure"
	registroMantenimientoInfra "github.com/vicpoo/API_recolecta/src/Mantenimiento/infrastructure"
	reporteMantenimientoGeneradoInfra "github.com/vicpoo/API_recolecta/src/Mantenimiento/infrastructure"
	tipoMantenimientoInfra "github.com/vicpoo/API_recolecta/src/Mantenimiento/infrastructure"

	anomaliaInfra "github.com/vicpoo/API_recolecta/src/Fallas/infrastructure"
	incidenciaInfra "github.com/vicpoo/API_recolecta/src/Fallas/infrastructure"
	reporteConductorInfra "github.com/vicpoo/API_recolecta/src/Fallas/infrastructure"
	reporteFallaCriticaInfra "github.com/vicpoo/API_recolecta/src/Fallas/infrastructure"
	seguimientoFallaCriticaInfra "github.com/vicpoo/API_recolecta/src/Fallas/infrastructure"

	notificacionInfra "github.com/vicpoo/API_recolecta/src/notificacion/infrastructure"

	coloniaHttp "github.com/vicpoo/API_recolecta/src/colonia/infrastructure/http"
	domicilioHttp "github.com/vicpoo/API_recolecta/src/domicilio/infrastructure/http"

	rolInfra "github.com/vicpoo/API_recolecta/src/rol/infrastructure"
	usuarioInfra "github.com/vicpoo/API_recolecta/src/usuario/infrastructure"

	alertaHttp "github.com/vicpoo/API_recolecta/src/alerta_usuario/infrastructure/http"
)

// @title           API Recolecta
// @version         1.0
// @description     API para gestión de recolección de residuos
// @host            localhost:8080
// @BasePath        /
func NewServer() *gin.Engine {
	server := gin.Default()
	server.Use(core.CORSMiddleware())
	server.Use(core.RateLimitMiddleware(100, time.Hour)) // Limita a 100 peticiones por hora por IP
	server.Use(core.SecurityHeadersMiddleware())
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return server
}

func RunServer(
	server *gin.Engine,

	tipoCamion *tipoCamionRoutes.TipoCamionRoutes,
	historial  *historialRoutes.HistorialAsignacionCamionRoutes,
	rutaCamion *rutaCamionRoutes.RutaCamionRoutes,

	camion          *camionRoutes.CamionRoutes,
	estadoCamion    *estadoCamionRoutes.EstadoCamionRoutes,
	ruta            *rutaRoutes.RutaRoutes,
	punto           *puntoRoutes.PuntoRecoleccionRoutes,
	rs              *rsRoutes.RellenoSanitarioRoutes,
	registroVaciado *registroVaciadoRoutes.RegistroVaciadoRoutes,

	alertaMantenimiento     *alertaMantenimientoInfra.AlertaMantenimientoRouter,
	registroMantenimiento   *registroMantenimientoInfra.RegistroMantenimientoRouter,
	reporteMantenimientoGen *reporteMantenimientoGeneradoInfra.ReporteMantenimientoGeneradoRouter,
	tipoMantenimiento       *tipoMantenimientoInfra.TipoMantenimientoRouter,

	anomalia                *anomaliaInfra.AnomaliaRouter,
	incidencia              *incidenciaInfra.IncidenciaRouter,
	reporteConductor        *reporteConductorInfra.ReporteConductorRouter,
	reporteFallaCritica     *reporteFallaCriticaInfra.ReporteFallaCriticaRouter,
	seguimientoFallaCritica *seguimientoFallaCriticaInfra.SeguimientoFallaCriticaRouter,

	notificacion *notificacionInfra.NotificacionRouter,

	coloniaController   *coloniaHttp.ColoniaController,
	domicilioController *domicilioHttp.DomicilioController,

	usuarioRoutes *usuarioInfra.UsuarioRoutes,
	rolRoutes     *rolInfra.RolRoutes,

	alertaController *alertaHttp.AlertaController,
) {
	tipoCamion.Run()
	historial.Run()
	rutaCamion.Run()

	camion.Run()
	estadoCamion.Run()
	ruta.Run()
	punto.Run()
	rs.Run()
	registroVaciado.Run()

	alertaMantenimiento.Run()
	registroMantenimiento.Run()
	reporteMantenimientoGen.Run()
	tipoMantenimiento.Run()

	anomalia.Run()
	incidencia.Run()
	reporteConductor.Run()
	reporteFallaCritica.Run()
	seguimientoFallaCritica.Run()

	notificacion.Run()

	coloniaController.RegisterRoutes(server)
	domicilioController.RegisterRoutes(server)

	usuarioRoutes.Register()
	rolRoutes.Register()

	// alerta_usuario — grupo /api para mantener consistencia con el resto
	apiGroup := server.Group("/api")
	alertaController.RegisterRoutes(apiGroup)

	server.Run(":8080")
}