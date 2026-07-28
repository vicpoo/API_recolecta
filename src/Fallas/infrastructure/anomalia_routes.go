// anomalia_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/core"
	alertaDomain "github.com/vicpoo/API_recolecta/src/alerta_usuario/domain"
)

type AnomaliaRouter struct {
	engine                   *gin.Engine
	alertaRepo               alertaDomain.AlertaUsuarioRepository
	modeloReportesURL        string
	clasificadorURL          string
	anomaliaCreadaWebhookURL string
}

func NewAnomaliaRouter(engine *gin.Engine, alertaRepo alertaDomain.AlertaUsuarioRepository, modeloReportesURL, clasificadorURL, anomaliaCreadaWebhookURL string) *AnomaliaRouter {
	return &AnomaliaRouter{
		engine:                   engine,
		alertaRepo:               alertaRepo,
		modeloReportesURL:        modeloReportesURL,
		clasificadorURL:          clasificadorURL,
		anomaliaCreadaWebhookURL: anomaliaCreadaWebhookURL,
	}
}

func (router *AnomaliaRouter) Run() {
	// Inicializar dependencias
	createController, getByIdController, updateController, deleteController,
		getAllController, getByPuntoIDController, getByChoferIDController,
		getByCamionIDController, getByRutaIDController, getByReferenciaIDController,
		getByEstadoController, getByTipoAnomaliaController, getByFechaRangeController,
		getMisAnomaliasController, pipelineRetryWorker := InitAnomaliaDependencies(router.alertaRepo, router.modeloReportesURL, router.clasificadorURL, router.anomaliaCreadaWebhookURL)

	// Red de seguridad del pipeline modelo_reportes -> clasificador_reportes:
	// corre durante toda la vida del proceso, igual que tracking_ws.Hub mas
	// abajo en dependencies.go. Ver pipeline_retry_worker.go.
	go pipelineRetryWorker.Run()

	// Rutas abiertas a cualquier usuario autenticado (ciudadano, conductor o
	// staff), sin RequireRole -- la autorizacion fina (quien puede hacer
	// que) se resuelve dentro del controller/use case, no aqui:
	//   - POST "/": crear un reporte propio.
	//   - DELETE "/:id": borrar un reporte propio (staff puede borrar
	//     cualquiera; conductor/ciudadano solo el suyo -- ver
	//     DeleteAnomaliaUseCase).
	//   - GET "/mis-reportes": listar los reportes propios (ciudadano o
	//     conductor, segun el JWT -- ver GetMisAnomaliasController). Gin
	//     resuelve el segmento estatico "mis-reportes" antes que el
	//     wildcard "/:id" del grupo de abajo, asi que no chocan aunque
	//     esten al mismo nivel de la ruta.
	// Los ciudadanos no tienen role_id en el esquema de roles de empleados
	// (su JWT trae role_id: 0, ver login_ciudadano.go) y CONDUCTOR (4)
	// tampoco es staff, asi que estas rutas no pueden ir en el grupo de
	// abajo (RequireRole las bloquearia).
	abiertoGroup := router.engine.Group("/api/anomalias")
	abiertoGroup.Use(core.JWTAuthMiddleware())
	{
		abiertoGroup.POST("/", createController.Run)
		abiertoGroup.DELETE("/:id", deleteController.Run)
		abiertoGroup.GET("/mis-reportes", getMisAnomaliasController.Run)
	}

	// Resto del CRUD: solo staff (ADMIN/SUPERVISOR/COORDINADOR)
	anomaliaGroup := router.engine.Group("/api/anomalias")
	anomaliaGroup.Use(core.JWTAuthMiddleware(), core.RequireRole(core.ADMIN, core.SUPERVISOR, core.COORDINADOR))
	{
		anomaliaGroup.GET("/:id", getByIdController.Run)
		anomaliaGroup.PUT("/:id", updateController.Run)
		anomaliaGroup.GET("/", getAllController.Run)

		// Rutas específicas
		anomaliaGroup.GET("/punto/:puntoId", getByPuntoIDController.Run)
		anomaliaGroup.GET("/chofer/:choferId", getByChoferIDController.Run)
		anomaliaGroup.GET("/camion/:camionId", getByCamionIDController.Run)
		anomaliaGroup.GET("/ruta/:rutaId", getByRutaIDController.Run)
		anomaliaGroup.GET("/referencia/:referenciaId", getByReferenciaIDController.Run)
		anomaliaGroup.GET("/estado", getByEstadoController.Run)        // Query param: ?estado=
		anomaliaGroup.GET("/tipo", getByTipoAnomaliaController.Run)    // Query param: ?tipo_anomalia=
		anomaliaGroup.GET("/por-fecha", getByFechaRangeController.Run) // Query params: ?fecha_inicio=&fecha_fin=
	}
}
