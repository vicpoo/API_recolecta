package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/vicpoo/API_recolecta/src/core"
	"go.uber.org/dig"

	historialUseCases    "github.com/vicpoo/API_recolecta/src/Camion/application"
	rutaCamionApp        "github.com/vicpoo/API_recolecta/src/Camion/application"
	tipoCamionUseCases   "github.com/vicpoo/API_recolecta/src/Camion/application"
	historialAdapters    "github.com/vicpoo/API_recolecta/src/Camion/infrastructure/adapters"
	rutaCamionAdapters   "github.com/vicpoo/API_recolecta/src/Camion/infrastructure/adapters"
	tipoCamionAdapters   "github.com/vicpoo/API_recolecta/src/Camion/infrastructure/adapters"
	historialControllers "github.com/vicpoo/API_recolecta/src/Camion/infrastructure/controllers"
	rutaCamionControllers "github.com/vicpoo/API_recolecta/src/Camion/infrastructure/controllers"
	tipoCamionControllers "github.com/vicpoo/API_recolecta/src/Camion/infrastructure/controllers"
	historialRoutes      "github.com/vicpoo/API_recolecta/src/Camion/infrastructure/routes"
	rutaCamionRoutes     "github.com/vicpoo/API_recolecta/src/Camion/infrastructure/routes"
	tipoCamionRoutes     "github.com/vicpoo/API_recolecta/src/Camion/infrastructure/routes"
	alertaMantenimientoInfra "github.com/vicpoo/API_recolecta/src/Mantenimiento/infrastructure"
	camionUseCases       "github.com/vicpoo/API_recolecta/src/Rutas/application"
	estadoCamionUseCases "github.com/vicpoo/API_recolecta/src/Rutas/application"
	puntoUseCases        "github.com/vicpoo/API_recolecta/src/Rutas/application"
	registroVaciadoApplication "github.com/vicpoo/API_recolecta/src/Rutas/application"
	rsApplication        "github.com/vicpoo/API_recolecta/src/Rutas/application"
	rutaUseCases         "github.com/vicpoo/API_recolecta/src/Rutas/application"
	camionAdapters       "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/adapters"
	estadoCamionAdapters "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/adapters"
	puntoAdapters        "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/adapters"
	registroVaciadoAdapters "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/adapters"
	rsAdapters           "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/adapters"
	rutaAdapters         "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/adapters"
	camionControllers    "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/controllers"
	estadoCamionControllers "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/controllers"
	puntoControllers     "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/controllers"
	registroVaciadoControllers "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/controllers"
	rsControllers        "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/controllers"
	rutaControllers      "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/controllers"
	camionRoutes         "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/routes"
	estadoCamionRoutes   "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/routes"
	puntoRoutes          "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/routes"
	registroVaciadoRoutesPkg "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/routes"
	rsRoutes             "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/routes"
	rutaRoutes           "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/routes"
	notificacionInfra    "github.com/vicpoo/API_recolecta/src/notificacion/infrastructure"
	anomalia             "github.com/vicpoo/API_recolecta/src/Fallas/infrastructure"
	incidencia           "github.com/vicpoo/API_recolecta/src/Fallas/infrastructure"
	reporteConductor     "github.com/vicpoo/API_recolecta/src/Fallas/infrastructure"
	reporteFallaCritica  "github.com/vicpoo/API_recolecta/src/Fallas/infrastructure"
	seguimientoFallaCritica "github.com/vicpoo/API_recolecta/src/Fallas/infrastructure"
	registroMantenimiento "github.com/vicpoo/API_recolecta/src/Mantenimiento/infrastructure"
	reporteMantenimientoGenerado "github.com/vicpoo/API_recolecta/src/Mantenimiento/infrastructure"
	tipoMantenimiento    "github.com/vicpoo/API_recolecta/src/Mantenimiento/infrastructure"
	coloniaApplication   "github.com/vicpoo/API_recolecta/src/colonia/application"
	coloniaHttp          "github.com/vicpoo/API_recolecta/src/colonia/infrastructure/http"
	coloniaPostgres      "github.com/vicpoo/API_recolecta/src/colonia/infrastructure/postgres"
	domicilioApplication "github.com/vicpoo/API_recolecta/src/domicilio/application"
	domicilioHttp        "github.com/vicpoo/API_recolecta/src/domicilio/infrastructure/http"
	domicilioPostgres    "github.com/vicpoo/API_recolecta/src/domicilio/infrastructure/postgres"
	_ "github.com/vicpoo/API_recolecta/src/notificacion/infrastructure"
	rolInfra             "github.com/vicpoo/API_recolecta/src/rol/infrastructure"
	usuarioInfra         "github.com/vicpoo/API_recolecta/src/usuario/infrastructure"
	usuarioController    "github.com/vicpoo/API_recolecta/src/usuario/infrastructure/controller"

	// ── alerta_usuario ──────────────────────────────────────────────────────
	alertaApp     "github.com/vicpoo/API_recolecta/src/alerta_usuario/application"
	alertaHttp    "github.com/vicpoo/API_recolecta/src/alerta_usuario/infrastructure/http"
	alertaPostgres "github.com/vicpoo/API_recolecta/src/alerta_usuario/infrastructure/postgres"
)

func InitDependencies() {
	if errEnv := godotenv.Load(); errEnv != nil {
		log.Fatal("error al cargar el .env")
	}
	container := dig.New()

	container.Provide(core.ConnectPostgres)
	container.Provide(NewServer)

	container.Provide(alertaMantenimientoInfra.NewAlertaMantenimientoRouter)
	container.Provide(notificacionInfra.NewNotificacionRouter)

	container.Provide(tipoCamionAdapters.NewPostgresTipoCamion)
	container.Provide(tipoCamionUseCases.NewSaveTipoCamionUseCase)
	container.Provide(tipoCamionUseCases.NewListAllTipoCamion)
	container.Provide(tipoCamionUseCases.NewGetTipoCamionByNameUseCase)
	container.Provide(tipoCamionUseCases.NewDeleteTipoCamionUseCase)
	container.Provide(tipoCamionControllers.NewCreateTipoCamionController)
	container.Provide(tipoCamionControllers.NewGetAllTipoCamionController)
	container.Provide(tipoCamionControllers.NewGetTipoCamionByNameController)
	container.Provide(tipoCamionControllers.NewDeleteTipoCamionController)
	container.Provide(tipoCamionRoutes.NewTipoCamionRoutes)

	container.Provide(camionAdapters.NewPostgresCamion)
	container.Provide(camionUseCases.NewSaveCamionUseCase)
	container.Provide(camionUseCases.NewListCamionUseCase)
	container.Provide(camionUseCases.NewUpdateCamionUseCase)
	container.Provide(camionUseCases.NewDeleteCamionUseCase)
	container.Provide(camionUseCases.NewGetCamionByIDUseCase)
	container.Provide(camionUseCases.NewGetCamionByPlacaUseCase)
	container.Provide(camionUseCases.NewGetCamionByModeloUseCase)
	container.Provide(camionControllers.NewCreateCamionController)
	container.Provide(camionControllers.NewGetAllCamionController)
	container.Provide(camionControllers.NewUpdateCamionController)
	container.Provide(camionControllers.NewDeleteCamionController)
	container.Provide(camionControllers.NewGetCamionByIDController)
	container.Provide(camionControllers.NewGetCamionByPlacaController)
	container.Provide(camionControllers.NewGetCamionByModeloController)
	container.Provide(camionRoutes.NewCamionRoutes)

	container.Provide(estadoCamionAdapters.NewPostgresEstadoCamion)
	container.Provide(estadoCamionUseCases.NewSaveEstadoCamionUseCase)
	container.Provide(estadoCamionUseCases.NewListAllEstadoCamionUseCase)
	container.Provide(estadoCamionUseCases.NewGetByIdEstadoCamionUseCase)
	container.Provide(estadoCamionUseCases.NewUpdateEstadoCamionUseCase)
	container.Provide(estadoCamionUseCases.NewDeleteEstadoCamionUseCase)
	container.Provide(estadoCamionControllers.NewCreateEstadoCamionController)
	container.Provide(estadoCamionControllers.NewGetAllEstadoCamionController)
	container.Provide(estadoCamionControllers.NewGetEstadoCamionByIdController)
	container.Provide(estadoCamionControllers.NewUpdateEstadoCamionController)
	container.Provide(estadoCamionControllers.NewDeleteEstadoCamionController)
	container.Provide(estadoCamionRoutes.NewEstadoCamionRoutes)

	container.Provide(historialAdapters.NewPostgresHistorialAsignacionCamion)
	container.Provide(historialUseCases.NewSaveHistorialAsignacionCamionUseCase)
	container.Provide(historialUseCases.NewListAllHistorialAsignacionCamionUseCase)
	container.Provide(historialUseCases.NewGetHistorialAsignacionCamionByIdUseCase)
	container.Provide(historialUseCases.NewUpdateHistorialAsignacionCamionUseCase)
	container.Provide(historialUseCases.NewDeleteHistorialAsignacionCamionUseCase)
	container.Provide(historialUseCases.NewGetHistorialByCamionUseCase)
	container.Provide(historialUseCases.NewGetHistorialByChoferUseCase)
	container.Provide(historialUseCases.NewGetActivoByCamionUseCase)
	container.Provide(historialUseCases.NewGetActivoByChoferUseCase)
	container.Provide(historialUseCases.NewDarDeBajaHistorialAsignacionUseCase)
	container.Provide(historialUseCases.NewCerrarAsignacionActivaCamionUseCase)
	container.Provide(historialUseCases.NewCerrarAsignacionActivaChoferUseCase)
	container.Provide(historialControllers.NewCreateHistorialAsignacionCamionController)
	container.Provide(historialControllers.NewGetAllHistorialAsignacionCamionController)
	container.Provide(historialControllers.NewGetHistorialAsignacionByIdController)
	container.Provide(historialControllers.NewUpdateHistorialAsignacionCamionController)
	container.Provide(historialControllers.NewDeleteHistorialAsignacionCamionController)
	container.Provide(historialControllers.NewGetHistorialByCamionController)
	container.Provide(historialControllers.NewGetHistorialByChoferController)
	container.Provide(historialControllers.NewGetActivoByCamionController)
	container.Provide(historialControllers.NewGetActivoByChoferController)
	container.Provide(historialControllers.NewDarDeBajaHistorialAsignacionController)
	container.Provide(historialControllers.NewCerrarAsignacionActivaCamionController)
	container.Provide(historialControllers.NewCerrarAsignacionActivaChoferController)
	container.Provide(historialRoutes.NewHistorialAsignacionCamionRoutes)

	container.Provide(rutaAdapters.NewPostgresRuta)
	container.Provide(rutaUseCases.NewCreateRutaUseCase)
	container.Provide(rutaUseCases.NewListAllRutaUseCase)
	container.Provide(rutaUseCases.NewGetRutaByIdUseCase)
	container.Provide(rutaUseCases.NewUpdateRutaUseCase)
	container.Provide(rutaUseCases.NewDeleteRutaUseCase)
	container.Provide(rutaUseCases.NewGetRutaActivasUseCase)
	container.Provide(rutaControllers.NewCreateRutaController)
	container.Provide(rutaControllers.NewGetAllRutaController)
	container.Provide(rutaControllers.NewGetRutaByIdController)
	container.Provide(rutaControllers.NewUpdateRutaController)
	container.Provide(rutaControllers.NewDeleteRutaController)
	container.Provide(rutaControllers.NewGetRutaActivasController)
	container.Provide(rutaRoutes.NewRutaRoutes)

	container.Provide(puntoAdapters.NewPostgresPuntoRecoleccion)
	container.Provide(puntoUseCases.NewSavePuntoRecoleccionUseCase)
	container.Provide(puntoUseCases.NewUpdatePuntoRecoleccionUseCase)
	container.Provide(puntoUseCases.NewListAllPuntoRecoleccionUseCase)
	container.Provide(puntoUseCases.NewGetPuntoRecoleccionByIdUseCase)
	container.Provide(puntoUseCases.NewGetPuntoRecoleccionByRutaUseCase)
	container.Provide(puntoUseCases.NewDeletePuntoRecoleccionUseCase)
	container.Provide(puntoControllers.NewCreatePuntoRecoleccionController)
	container.Provide(puntoControllers.NewUpdatePuntoRecoleccionController)
	container.Provide(puntoControllers.NewGetAllPuntoRecoleccionController)
	container.Provide(puntoControllers.NewGetPuntoRecoleccionByIdController)
	container.Provide(puntoControllers.NewGetPuntoRecoleccionByRutaController)
	container.Provide(puntoControllers.NewDeletePuntoRecoleccionController)
	container.Provide(puntoRoutes.NewPuntoRecoleccionRoutes)

	container.Provide(rsAdapters.NewPostgresRellenoSanitario)
	container.Provide(rsApplication.NewSaveRellenoSanitarioUseCase)
	container.Provide(rsApplication.NewUpdateRellenoSanitarioUseCase)
	container.Provide(rsApplication.NewListRellenoSanitarioUseCase)
	container.Provide(rsApplication.NewGetRellenoSanitarioByIdUseCase)
	container.Provide(rsApplication.NewDeleteRellenoSanitarioUseCase)
	container.Provide(rsApplication.NewGetRellenoSanitarioByNombreUseCase)
	container.Provide(rsApplication.NewExistsRellenoSanitarioByIdUseCase)
	container.Provide(rsControllers.NewCreateRellenoSanitarioController)
	container.Provide(rsControllers.NewUpdateRellenoSanitarioController)
	container.Provide(rsControllers.NewGetAllRellenoSanitarioController)
	container.Provide(rsControllers.NewGetRellenoSanitarioByIDController)
	container.Provide(rsControllers.NewDeleteRellenoSanitarioController)
	container.Provide(rsControllers.NewGetRellenoSanitarioByNombreController)
	container.Provide(rsControllers.NewExistsRellenoSanitarioByIdController)
	container.Provide(rsRoutes.NewRellenoSanitarioRoutes)

	container.Provide(rutaCamionAdapters.NewPostgresRutaCamion)
	container.Provide(rutaCamionApp.NewSaveRutaCamionUseCase)
	container.Provide(rutaCamionApp.NewUpdateRutaCamionUseCase)
	container.Provide(rutaCamionApp.NewListAllRutaCamionUseCase)
	container.Provide(rutaCamionApp.NewGetRutaCamionByIDUseCase)
	container.Provide(rutaCamionApp.NewGetRutaCamionByCamionIDUseCase)
	container.Provide(rutaCamionApp.NewGetRutaCamionByRutaIDUseCase)
	container.Provide(rutaCamionApp.NewExistsRutaCamionByIDUseCase)
	container.Provide(rutaCamionApp.NewDeleteRutaCamionUseCase)
	container.Provide(rutaCamionControllers.NewCreateRutaCamionController)
	container.Provide(rutaCamionControllers.NewUpdateRutaCamionController)
	container.Provide(rutaCamionControllers.NewGetAllRutaCamionController)
	container.Provide(rutaCamionControllers.NewGetRutaCamionByIDController)
	container.Provide(rutaCamionControllers.NewGetRutaCamionByCamionIDController)
	container.Provide(rutaCamionControllers.NewGetRutaCamionByRutaIDController)
	container.Provide(rutaCamionControllers.NewExistsRutaCamionByIDController)
	container.Provide(rutaCamionControllers.NewDeleteRutaCamionController)
	container.Provide(rutaCamionRoutes.NewRutaCamionRoutes)

	container.Provide(registroVaciadoAdapters.NewPostgresRegistroVaciado)
	container.Provide(registroVaciadoApplication.NewCreateRegistroVaciadoUseCase)
	container.Provide(registroVaciadoApplication.NewListAllRegistroVaciadoUseCase)
	container.Provide(registroVaciadoApplication.NewGetRegistroVaciadoByIDUseCase)
	container.Provide(registroVaciadoApplication.NewGetRegistroVaciadoByRellenoIDUseCase)
	container.Provide(registroVaciadoApplication.NewGetRegistroVaciadoByRutaCamionIDUseCase)
	container.Provide(registroVaciadoApplication.NewExistsRegistroVaciadoUseCase)
	container.Provide(registroVaciadoApplication.NewDeleteRegistroVaciadoUseCase)
	container.Provide(registroVaciadoControllers.NewCreateRegistroVaciadoController)
	container.Provide(registroVaciadoControllers.NewGetAllRegistroVaciadoController)
	container.Provide(registroVaciadoControllers.NewGetRegistroVaciadoByIDController)
	container.Provide(registroVaciadoControllers.NewGetRegistroVaciadoByRellenoIDController)
	container.Provide(registroVaciadoControllers.NewGetRegistroVaciadoByRutaCamionIDController)
	container.Provide(registroVaciadoControllers.NewExistsRegistroVaciadoController)
	container.Provide(registroVaciadoControllers.NewDeleteRegistroVaciadoController)
	container.Provide(registroVaciadoRoutesPkg.NewRegistroVaciadoRoutes)

	container.Provide(coloniaPostgres.NewColoniaRepository)
	container.Provide(coloniaApplication.NewCreateColonia)
	container.Provide(coloniaApplication.NewGetColonia)
	container.Provide(coloniaApplication.NewListColonias)
	container.Provide(coloniaApplication.NewUpdateColonia)
	container.Provide(coloniaApplication.NewDeleteColonia)
	container.Provide(coloniaHttp.NewColoniaController)

	container.Provide(domicilioPostgres.NewDomicilioRepository)
	container.Provide(domicilioApplication.NewCreateDomicilio)
	container.Provide(domicilioApplication.NewGetDomicilio)
	container.Provide(domicilioApplication.NewUpdateDomicilio)
	container.Provide(domicilioApplication.NewDeleteDomicilio)
	container.Provide(domicilioHttp.NewDomicilioController)

	container.Provide(usuarioInfra.NewUsuarioDependencies)
	container.Provide(func(d *usuarioInfra.UsuarioDependencies) *usuarioController.AddUsersController    { return d.Create })
	container.Provide(func(d *usuarioInfra.UsuarioDependencies) *usuarioController.DeleteUsersController { return d.Delete })
	container.Provide(func(d *usuarioInfra.UsuarioDependencies) *usuarioController.ViewOneUsersController { return d.Get })
	container.Provide(func(d *usuarioInfra.UsuarioDependencies) *usuarioController.ViewAllUsersController { return d.List })
	container.Provide(func(d *usuarioInfra.UsuarioDependencies) *usuarioController.LoginUsersController  { return d.Login })
	container.Provide(usuarioInfra.NewUsuarioRoutes)

	container.Provide(rolInfra.NewRolDependencies)
	container.Provide(rolInfra.NewRolRoutes)

	container.Provide(anomalia.NewAnomaliaRouter)
	container.Provide(incidencia.NewIncidenciaRouter)
	container.Provide(reporteConductor.NewReporteConductorRouter)
	container.Provide(registroMantenimiento.NewRegistroMantenimientoRouter)
	container.Provide(reporteFallaCritica.NewReporteFallaCriticaRouter)
	container.Provide(reporteMantenimientoGenerado.NewReporteMantenimientoGeneradoRouter)
	container.Provide(seguimientoFallaCritica.NewSeguimientoFallaCriticaRouter)
	container.Provide(tipoMantenimiento.NewTipoMantenimientoRouter)

	// ── alerta_usuario ──────────────────────────────────────────────────────
	container.Provide(alertaPostgres.NewPostgresAlertaRepository)
	container.Provide(alertaApp.NewCreateAlerta)
	container.Provide(alertaApp.NewListMisAlertas)
	container.Provide(alertaApp.NewMarcarLeida)
	container.Provide(alertaHttp.NewAlertaController)

	if err := container.Invoke(RunServer); err != nil {
		log.Fatal(err)
	}
}