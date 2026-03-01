package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vicpoo/API_recolecta/src/core"

	// Agregar esta línea
	alertaMantenimientoInfra "github.com/vicpoo/API_recolecta/src/alerta_mantenimiento/infrastructure"
	notificacionInfra "github.com/vicpoo/API_recolecta/src/notificacion/infrastructure"
	camionUseCases "github.com/vicpoo/API_recolecta/src/Rutas/application"
	camionAdapters "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/adapters"
	camionControllers "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/controllers"
	camionRoutes "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/routes"
	estadoCamionUseCases "github.com/vicpoo/API_recolecta/src/Rutas/application"
	estadoCamionAdapters "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/adapters"
	estadoCamionControllers "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/controllers"
	estadoCamionRoutes "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/routes"
	historialUseCases "github.com/vicpoo/API_recolecta/src/Camion/application"
	historialAdapters "github.com/vicpoo/API_recolecta/src/Camion/infrastructure/adapters"
	historialControllers "github.com/vicpoo/API_recolecta/src/Camion/infrastructure/controllers"
	historialRoutes "github.com/vicpoo/API_recolecta/src/Camion/infrastructure/routes"
	puntoUseCases "github.com/vicpoo/API_recolecta/src/Rutas/application"
	puntoAdapters "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/adapters"
	puntoControllers "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/controllers"
	puntoRoutes "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/routes"
	rsApplication "github.com/vicpoo/API_recolecta/src/Rutas/application"
	rsAdapters "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/adapters"
	rsControllers "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/controllers"
	rsRoutes "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/routes"
	rutaUseCases "github.com/vicpoo/API_recolecta/src/Rutas/application"
	rutaAdapters "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/adapters"
	rutaControllers "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/controllers"
	rutaRoutes "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/routes"
	rutaCamionApp "github.com/vicpoo/API_recolecta/src/Camion/application"
	rutaCamionAdapters "github.com/vicpoo/API_recolecta/src/Camion/infrastructure/adapters"
	rutaCamionControllers "github.com/vicpoo/API_recolecta/src/Camion/infrastructure/controllers"
	rutaCamionRoutes "github.com/vicpoo/API_recolecta/src/Camion/infrastructure/routes"
	tipoCamionUseCases "github.com/vicpoo/API_recolecta/src/Camion/application"
	tipoCamionAdapters "github.com/vicpoo/API_recolecta/src/Camion/infrastructure/adapters"
	tipoCamionControllers "github.com/vicpoo/API_recolecta/src/Camion/infrastructure/controllers"
	tipoCamionRoutes "github.com/vicpoo/API_recolecta/src/Camion/infrastructure/routes"
	registroVaciadoAdapters "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/adapters"
	registroVaciadoApplication "github.com/vicpoo/API_recolecta/src/Rutas/application"
	registroVaciadoControllers "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/controllers"
	registroVaciadoRoutesPkg "github.com/vicpoo/API_recolecta/src/Rutas/infraestructure/routes"
	anomalia "github.com/vicpoo/API_recolecta/src/anomalia/infrastructure"
	incidencia "github.com/vicpoo/API_recolecta/src/incidencia/infrastructure"
	_ "github.com/vicpoo/API_recolecta/src/notificacion/infrastructure"
	reporteConductor "github.com/vicpoo/API_recolecta/src/reporte_conductor/infrastructure"
	registroMantenimiento "github.com/vicpoo/API_recolecta/src/registro_mantenimiento/infrastructure"
	reporteFallaCritica "github.com/vicpoo/API_recolecta/src/reporte_falla_critica/infrastructure"
	reporteMantenimientoGenerado "github.com/vicpoo/API_recolecta/src/reporte_mantenimiento_generado/infrastructure"
	seguimientoFallaCritica "github.com/vicpoo/API_recolecta/src/seguimiento_falla_critica/infrastructure"
	tipoMantenimiento "github.com/vicpoo/API_recolecta/src/tipo_mantenimiento/infrastructure"
	domicilioApplication "github.com/vicpoo/API_recolecta/src/domicilio/application"
	domicilioHttp "github.com/vicpoo/API_recolecta/src/domicilio/infrastructure/http"
	domicilioPostgres "github.com/vicpoo/API_recolecta/src/domicilio/infrastructure/postgres"
	coloniaPostgres "github.com/vicpoo/API_recolecta/src/colonia/infrastructure/postgres"
	coloniaApplication "github.com/vicpoo/API_recolecta/src/colonia/application"
	coloniaHttp "github.com/vicpoo/API_recolecta/src/colonia/infrastructure/http"
	rolInfra "github.com/vicpoo/API_recolecta/src/rol/infrastructure"
	usuarioInfra "github.com/vicpoo/API_recolecta/src/usuario/infrastructure"
)

// archivo para hacer las instancias de los controllers, casos de uso y repositories, etc.
func InitDependencies() {
	if errEnv := godotenv.Load(); errEnv != nil {
		log.Fatal("error al cargar el .env")
	}

	engine := gin.Default()
	engine.Use(core.CORSMiddleware())

	db := core.GetBD()

	// ================================
	// ALERTAS MANTENIMIENTO
	// ================================
	alertaMantenimientoRoutes := alertaMantenimientoInfra.NewAlertaMantenimientoRouter(engine)
	alertaMantenimientoRoutes.Run()

	// =================================
	// NOTIFICACIONES
	// =================================
	notificacionRoutes := notificacionInfra.NewNotificacionRouter(engine)
	notificacionRoutes.Run()

	// tipo camion
	tipoCamionRepository := tipoCamionAdapters.NewPostgresTipoCamion()
	saveTipoCamionUc := tipoCamionUseCases.NewSaveTipoCamionUseCase(tipoCamionRepository)
	listAllTipoCamionUc := tipoCamionUseCases.NewListAllTipoCamion(tipoCamionRepository)
	getTipoCamionUc := tipoCamionUseCases.NewGetTipoCamionByNameUseCase(tipoCamionRepository)
	deleteTipoCamionByIdUc := tipoCamionUseCases.NewDeleteTipoCamionUseCase(tipoCamionRepository)

	createTipoCamionCtr := tipoCamionControllers.NewCreateTipoCamionController(saveTipoCamionUc)
	getAllTipoCamionCtr := tipoCamionControllers.NewGetAllTipoCamionController(listAllTipoCamionUc)
	getTipoCamionByNameCtr := tipoCamionControllers.NewGetTipoCamionByNameController(getTipoCamionUc)
	deleteTipoCamionByIdCtr := tipoCamionControllers.NewDeleteTipoCamionController(deleteTipoCamionByIdUc)

	tipoCamionRoutes := tipoCamionRoutes.NewTipoCamionRoutes(
		engine,
		createTipoCamionCtr,
		getAllTipoCamionCtr,
		getTipoCamionByNameCtr,
		deleteTipoCamionByIdCtr,
	)
	tipoCamionRoutes.Run()

	// camion
	camionRepository := camionAdapters.NewPostgresCamion()
	saveCamionUc := camionUseCases.NewSaveCamionUseCase(camionRepository)
	listAllCamionUc := camionUseCases.NewListCamionUseCase(camionRepository)
	updateCamionUc := camionUseCases.NewUpdateCamionUseCase(camionRepository)
	deleteCamionByIdUc := camionUseCases.NewDeleteCamionUseCase(camionRepository)
	getCamionByIdUc := camionUseCases.NewGetCamionByIDUseCase(camionRepository)
	getCamionByPlacaUc := camionUseCases.NewGetCamionByPlacaUseCase(camionRepository)
	getCamionByModeloUc := camionUseCases.NewGetCamionByModeloUseCase(camionRepository)

	createCamionCtr := camionControllers.NewCreateCamionController(saveCamionUc)
	getAllCamionCtr := camionControllers.NewGetAllCamionController(listAllCamionUc)
	updateCamionCtr := camionControllers.NewUpdateCamionController(updateCamionUc)
	deleteCamionByIdCtr := camionControllers.NewDeleteCamionController(deleteCamionByIdUc)
	getCamionByIdCtr := camionControllers.NewGetCamionByIDController(getCamionByIdUc)
	getCamionByPlacaCtr := camionControllers.NewGetCamionByPlacaController(getCamionByPlacaUc)
	getCamionByModeloCtr := camionControllers.NewGetCamionByModeloController(getCamionByModeloUc)

	camionRoutes := camionRoutes.NewCamionRoutes(
		engine, createCamionCtr,
		getAllCamionCtr,
		getCamionByIdCtr,
		updateCamionCtr,
		deleteCamionByIdCtr,
		getCamionByPlacaCtr,
		getCamionByModeloCtr,
	)
	camionRoutes.Run()

	// estado camion
	estadoCamionRepository := estadoCamionAdapters.NewPostgresEstadoCamion()

	saveEstadoCamionUc := estadoCamionUseCases.NewSaveEstadoCamionUseCase(estadoCamionRepository)
	listEstadoCamionUc := estadoCamionUseCases.NewListAllEstadoCamionUseCase(estadoCamionRepository)
	getEstadoCamionByIdUc := estadoCamionUseCases.NewGetByIdEstadoCamionUseCase(estadoCamionRepository)
	updateEstadoCamionUc := estadoCamionUseCases.NewUpdateEstadoCamionUseCase(estadoCamionRepository)
	deleteEstadoCamionUc := estadoCamionUseCases.NewDeleteEstadoCamionUseCase(estadoCamionRepository)

	createEstadoCamionCtr := estadoCamionControllers.NewCreateEstadoCamionController(saveEstadoCamionUc)
	getAllEstadoCamionCtr := estadoCamionControllers.NewGetAllEstadoCamionController(listEstadoCamionUc)
	getEstadoCamionByIdCtr := estadoCamionControllers.NewGetEstadoCamionByIdController(getEstadoCamionByIdUc)
	updateEstadoCamionCtr := estadoCamionControllers.NewUpdateEstadoCamionController(&updateEstadoCamionUc)
	deleteEstadoCamionCtr := estadoCamionControllers.NewDeleteEstadoCamionController(deleteEstadoCamionUc)

	estadoCamionRoutes := estadoCamionRoutes.NewEstadoCamionRoutes(
		engine,
		createEstadoCamionCtr,
		getAllEstadoCamionCtr,
		getEstadoCamionByIdCtr,
		deleteEstadoCamionCtr,
		updateEstadoCamionCtr,
	)

	estadoCamionRoutes.Run()

	// ================================
	// HISTORIAL ASIGNACION CAMION
	// ================================
	historialRepository := historialAdapters.NewPostgresHistorialAsignacionCamion()

	createHistorialUC := historialUseCases.NewSaveHistorialAsignacionCamionUseCase(historialRepository)
	getAllHistorialUC := historialUseCases.NewListAllHistorialAsignacionCamionUseCase(historialRepository)
	getHistorialByIdUC := historialUseCases.NewGetHistorialAsignacionCamionByIdUseCase(historialRepository)
	updateHistorialUC := historialUseCases.NewUpdateHistorialAsignacionCamionUseCase(historialRepository)
	deleteHistorialUC := historialUseCases.NewDeleteHistorialAsignacionCamionUseCase(historialRepository)

	getByCamionUC := historialUseCases.NewGetHistorialByCamionUseCase(historialRepository)
	getByChoferUC := historialUseCases.NewGetHistorialByChoferUseCase(historialRepository)
	getActivoByCamionUC := historialUseCases.NewGetActivoByCamionUseCase(historialRepository)
	getActivoByChoferUC := historialUseCases.NewGetActivoByChoferUseCase(historialRepository)

	darDeBajaUC := historialUseCases.NewDarDeBajaHistorialAsignacionUseCase(historialRepository)
	cerrarCamionUC := historialUseCases.NewCerrarAsignacionActivaCamionUseCase(historialRepository)
	cerrarChoferUC := historialUseCases.NewCerrarAsignacionActivaChoferUseCase(historialRepository)

	createHistorialCtr := historialControllers.NewCreateHistorialAsignacionCamionController(createHistorialUC)
	getAllHistorialCtr := historialControllers.NewGetAllHistorialAsignacionCamionController(getAllHistorialUC)
	getHistorialByIdCtr := historialControllers.NewGetHistorialAsignacionByIdController(getHistorialByIdUC)
	updateHistorialCtr := historialControllers.NewUpdateHistorialAsignacionCamionController(updateHistorialUC)
	deleteHistorialCtr := historialControllers.NewDeleteHistorialAsignacionCamionController(deleteHistorialUC)

	getByCamionCtr := historialControllers.NewGetHistorialByCamionController(getByCamionUC)
	getByChoferCtr := historialControllers.NewGetHistorialByChoferController(getByChoferUC)
	getActivoByCamionCtr := historialControllers.NewGetActivoByCamionController(getActivoByCamionUC)
	getActivoByChoferCtr := historialControllers.NewGetActivoByChoferController(getActivoByChoferUC)

	darDeBajaCtr := historialControllers.NewDarDeBajaHistorialAsignacionController(darDeBajaUC)
	cerrarCamionCtr := historialControllers.NewCerrarAsignacionActivaCamionController(cerrarCamionUC)
	cerrarChoferCtr := historialControllers.NewCerrarAsignacionActivaChoferController(cerrarChoferUC)

	historialRoutes := historialRoutes.NewHistorialAsignacionCamionRoutes(
		engine,
		createHistorialCtr,
		getAllHistorialCtr,
		getHistorialByIdCtr,
		updateHistorialCtr,
		deleteHistorialCtr,
		getByCamionCtr,
		getByChoferCtr,
		getActivoByCamionCtr,
		getActivoByChoferCtr,
		darDeBajaCtr,
		cerrarCamionCtr,
		cerrarChoferCtr,
	)

	historialRoutes.Run()

	// ================================
	// RUTA
	// ================================

	rutaRepository := rutaAdapters.NewPostgresRuta()

	createRutaUc := rutaUseCases.NewCreateRutaUseCase(rutaRepository)
	getAllRutaUc := rutaUseCases.NewListAllRutaUseCase(rutaRepository)
	getRutaByIdUc := rutaUseCases.NewGetRutaByIdUseCase(rutaRepository)
	updateRutaUc := rutaUseCases.NewUpdateRutaUseCase(rutaRepository)
	deleteRutaUc := rutaUseCases.NewDeleteRutaUseCase(rutaRepository)
	getRutasActivasUc := rutaUseCases.NewGetRutaActivasUseCase(rutaRepository)

	createRutaCtr := rutaControllers.NewCreateRutaController(createRutaUc)
	getAllRutaCtr := rutaControllers.NewGetAllRutaController(getAllRutaUc)
	getRutaByIdCtr := rutaControllers.NewGetRutaByIdController(getRutaByIdUc)
	updateRutaCtr := rutaControllers.NewUpdateRutaController(updateRutaUc)
	deleteRutaCtr := rutaControllers.NewDeleteRutaController(deleteRutaUc)
	getRutasActivasCtr := rutaControllers.NewGetRutaActivasController(getRutasActivasUc)

	rutaRoutes := rutaRoutes.NewRutaRoutes(
		engine,
		createRutaCtr,
		getAllRutaCtr,
		getRutaByIdCtr,
		updateRutaCtr,
		deleteRutaCtr,
		getRutasActivasCtr,
	)

	rutaRoutes.Run()

	puntoRepository := puntoAdapters.NewPostgresPuntoRecoleccion()

	createPuntoUC := puntoUseCases.NewSavePuntoRecoleccionUseCase(puntoRepository)
	updatePuntoUC := puntoUseCases.NewUpdatePuntoRecoleccionUseCase(puntoRepository)
	getAllPuntoUC := puntoUseCases.NewListAllPuntoRecoleccionUseCase(puntoRepository)
	getPuntoByIdUC := puntoUseCases.NewGetPuntoRecoleccionByIdUseCase(puntoRepository)
	getPuntoByRutaUC := puntoUseCases.NewGetPuntoRecoleccionByRutaUseCase(puntoRepository)
	deletePuntoUC := puntoUseCases.NewDeletePuntoRecoleccionUseCase(puntoRepository)

	createPuntoCTR := puntoControllers.NewCreatePuntoRecoleccionController(createPuntoUC)
	updatePuntoCTR := puntoControllers.NewUpdatePuntoRecoleccionController(updatePuntoUC)
	getAllPuntoCTR := puntoControllers.NewGetAllPuntoRecoleccionController(getAllPuntoUC)
	getPuntoByIdCTR := puntoControllers.NewGetPuntoRecoleccionByIdController(getPuntoByIdUC)
	getPuntoByRutaCTR := puntoControllers.NewGetPuntoRecoleccionByRutaController(getPuntoByRutaUC)
	deletePuntoCTR := puntoControllers.NewDeletePuntoRecoleccionController(deletePuntoUC)

	puntoRoutes := puntoRoutes.NewPuntoRecoleccionRoutes(
		engine,
		createPuntoCTR,
		getAllPuntoCTR,
		getPuntoByIdCTR,
		getPuntoByRutaCTR,
		updatePuntoCTR,
		deletePuntoCTR,
	)

	puntoRoutes.Run()

	rellenoRepo := rsAdapters.NewPostgresRellenoSanitario()

	createRellenoUC := rsApplication.NewSaveRellenoSanitarioUseCase(rellenoRepo)
	updateRellenoUC := rsApplication.NewUpdateRellenoSanitarioUseCase(rellenoRepo)
	getAllRellenoUC := rsApplication.NewListRellenoSanitarioUseCase(rellenoRepo)
	getRellenoByIDUC := rsApplication.NewGetRellenoSanitarioByIdUseCase(rellenoRepo)
	deleteRellenoUC := rsApplication.NewDeleteRellenoSanitarioUseCase(rellenoRepo)
	getRellenoByNombreUC := rsApplication.NewGetRellenoSanitarioByNombreUseCase(rellenoRepo)
	existsRellenoUC := rsApplication.NewExistsRellenoSanitarioByIdUseCase(rellenoRepo)

	createRellenoController := rsControllers.NewCreateRellenoSanitarioController(createRellenoUC)
	updateRellenoController := rsControllers.NewUpdateRellenoSanitarioController(updateRellenoUC)
	getAllRellenoController := rsControllers.NewGetAllRellenoSanitarioController(getAllRellenoUC)
	getRellenoByIDController := rsControllers.NewGetRellenoSanitarioByIDController(getRellenoByIDUC)
	deleteRellenoController := rsControllers.NewDeleteRellenoSanitarioController(deleteRellenoUC)
	getRellenoByNombreController := rsControllers.NewGetRellenoSanitarioByNombreController(getRellenoByNombreUC)
	existsRellenoController := rsControllers.NewExistsRellenoSanitarioByIdController(existsRellenoUC)

	rellenoRoutes := rsRoutes.NewRellenoSanitarioRoutes(
		engine,
		createRellenoController,
		getAllRellenoController,
		getRellenoByIDController,
		updateRellenoController,
		deleteRellenoController,
		getRellenoByNombreController,
		existsRellenoController,
	)

	rellenoRoutes.Run()

	repository := rutaCamionAdapters.NewPostgresRutaCamion()

	// ===============================
	// USE CASES
	// ===============================
	createRutaCamionUC := rutaCamionApp.NewSaveRutaCamionUseCase(repository)
	updateRutaCamionUC := rutaCamionApp.NewUpdateRutaCamionUseCase(repository)
	getAllRutaCamionUC := rutaCamionApp.NewListAllRutaCamionUseCase(repository)
	getRutaCamionByIDUC := rutaCamionApp.NewGetRutaCamionByIDUseCase(repository)
	getRutaCamionByCamionIDUC := rutaCamionApp.NewGetRutaCamionByCamionIDUseCase(repository)
	getRutaCamionByRutaIDUC := rutaCamionApp.NewGetRutaCamionByRutaIDUseCase(repository)
	existsRutaCamionUC := rutaCamionApp.NewExistsRutaCamionByIDUseCase(repository)
	deleteRutaCamionUC := rutaCamionApp.NewDeleteRutaCamionUseCase(repository)

	// ===============================
	// CONTROLLERS
	// ===============================
	createRutaCamionController :=
		rutaCamionControllers.NewCreateRutaCamionController(createRutaCamionUC)

	updateRutaCamionController :=
		rutaCamionControllers.NewUpdateRutaCamionController(updateRutaCamionUC)

	getAllRutaCamionController :=
		rutaCamionControllers.NewGetAllRutaCamionController(getAllRutaCamionUC)

	getRutaCamionByIDController :=
		rutaCamionControllers.NewGetRutaCamionByIDController(getRutaCamionByIDUC)

	getRutaCamionByCamionIDController :=
		rutaCamionControllers.NewGetRutaCamionByCamionIDController(getRutaCamionByCamionIDUC)

	getRutaCamionByRutaIDController :=
		rutaCamionControllers.NewGetRutaCamionByRutaIDController(getRutaCamionByRutaIDUC)

	existsRutaCamionController :=
		rutaCamionControllers.NewExistsRutaCamionByIDController(existsRutaCamionUC)

	deleteRutaCamionController :=
		rutaCamionControllers.NewDeleteRutaCamionController(deleteRutaCamionUC)

	rutaCamionRoutes := rutaCamionRoutes.NewRutaCamionRoutes(
		engine,
		createRutaCamionController,
		getAllRutaCamionController,
		getRutaCamionByIDController,
		getRutaCamionByCamionIDController,
		getRutaCamionByRutaIDController,
		existsRutaCamionController,
		updateRutaCamionController,
		deleteRutaCamionController,
	)

	rutaCamionRoutes.Run()

	// ===============================
	// REGISTRO VACIADO
	// ===============================

	// Repository
	registroVaciadoRepository := registroVaciadoAdapters.NewPostgresRegistroVaciado()

	// Use Cases - CORREGIDOS: Los nombres deben coincidir con lo que esperan los controllers
	createRegistroVaciadoUC := registroVaciadoApplication.NewCreateRegistroVaciadoUseCase(registroVaciadoRepository)
    getAllRegistroVaciadoUC := registroVaciadoApplication.NewListAllRegistroVaciadoUseCase(registroVaciadoRepository)
	getRegistroVaciadoByIDUC := registroVaciadoApplication.NewGetRegistroVaciadoByIDUseCase(registroVaciadoRepository)
	getRegistroVaciadoByRellenoIDUC := registroVaciadoApplication.NewGetRegistroVaciadoByRellenoIDUseCase(registroVaciadoRepository)
	getRegistroVaciadoByRutaCamionIDUC := registroVaciadoApplication.NewGetRegistroVaciadoByRutaCamionIDUseCase(registroVaciadoRepository)
	existsRegistroVaciadoUC := registroVaciadoApplication.NewExistsRegistroVaciadoUseCase(registroVaciadoRepository)
	deleteRegistroVaciadoUC := registroVaciadoApplication.NewDeleteRegistroVaciadoUseCase(registroVaciadoRepository)

	// ===============================
	// CONTROLLERS
	// ===============================
	createRegistroVaciadoController := registroVaciadoControllers.NewCreateRegistroVaciadoController(createRegistroVaciadoUC)
	getAllRegistroVaciadoController := registroVaciadoControllers.NewGetAllRegistroVaciadoController(getAllRegistroVaciadoUC)
	getRegistroVaciadoByIDController := registroVaciadoControllers.NewGetRegistroVaciadoByIDController(getRegistroVaciadoByIDUC)
	getRegistroVaciadoByRellenoIDController := registroVaciadoControllers.NewGetRegistroVaciadoByRellenoIDController(getRegistroVaciadoByRellenoIDUC)
	getRegistroVaciadoByRutaCamionIDController := registroVaciadoControllers.NewGetRegistroVaciadoByRutaCamionIDController(getRegistroVaciadoByRutaCamionIDUC)
	existsRegistroVaciadoController := registroVaciadoControllers.NewExistsRegistroVaciadoController(existsRegistroVaciadoUC)
	deleteRegistroVaciadoController := registroVaciadoControllers.NewDeleteRegistroVaciadoController(deleteRegistroVaciadoUC)

	// ===============================
	// ROUTES
	// ===============================
	registroVaciadoRoutes := registroVaciadoRoutesPkg.NewRegistroVaciadoRoutes(
		engine,
		createRegistroVaciadoController,
		getAllRegistroVaciadoController,
		getRegistroVaciadoByIDController,
		getRegistroVaciadoByRellenoIDController,
		getRegistroVaciadoByRutaCamionIDController,
		existsRegistroVaciadoController,
		deleteRegistroVaciadoController,
	)

	registroVaciadoRoutes.Run()

	// ===============================
	// COLONIA
	// ===============================

	coloniaRepository := coloniaPostgres.NewColoniaRepository(core.GetBD())

	createColoniaUC := coloniaApplication.NewCreateColonia(coloniaRepository)
	getColoniaUC := coloniaApplication.NewGetColonia(coloniaRepository)
	listColoniasUC := coloniaApplication.NewListColonias(coloniaRepository)
	updateColoniaUC := coloniaApplication.NewUpdateColonia(coloniaRepository)
	deleteColoniaUC := coloniaApplication.NewDeleteColonia(coloniaRepository)

	coloniaController := coloniaHttp.NewColoniaController(
		createColoniaUC,
		getColoniaUC,
		listColoniasUC,
		updateColoniaUC,
		deleteColoniaUC,
	)

	coloniaController.RegisterRoutes(engine)

	// ===============================
	// DOMICILIO
	// ===============================

	domicilioRepository := domicilioPostgres.NewDomicilioRepository(core.GetBD())

	createDomicilioUC := domicilioApplication.NewCreateDomicilio(domicilioRepository)
	getDomicilioUC := domicilioApplication.NewGetDomicilio(domicilioRepository)
	updateDomicilioUC := domicilioApplication.NewUpdateDomicilio(domicilioRepository)
	deleteDomicilioUC := domicilioApplication.NewDeleteDomicilio(domicilioRepository)

	domicilioController := domicilioHttp.NewDomicilioController(
		createDomicilioUC,
		getDomicilioUC,
		updateDomicilioUC,
		deleteDomicilioUC,
	)

	domicilioController.RegisterRoutes(engine)

	usuarioDeps := usuarioInfra.NewUsuarioDependencies(db)
	usuarioInfra.RegisterUsuarioRoutes(engine, usuarioDeps)

	rolController := rolInfra.NewRolDependencies(db)
	rolInfra.RegisterRolRoutes(engine, rolController)

	anomaliaRoutes := anomalia.NewAnomaliaRouter(engine)
	anomaliaRoutes.Run()

	incidenciaRoutes := incidencia.NewIncidenciaRouter(engine)
	incidenciaRoutes.Run()

	reporteConductorRoutes := reporteConductor.NewReporteConductorRouter(engine)
	reporteConductorRoutes.Run()

	registroMantenimientoRoutes := registroMantenimiento.NewRegistroMantenimientoRouter(engine)
	registroMantenimientoRoutes.Run()

	reporteFallaCriticaRoutes := reporteFallaCritica.NewReporteFallaCriticaRouter(engine)
	reporteFallaCriticaRoutes.Run()

	reporteMantenimientoGeneradoRoutes := reporteMantenimientoGenerado.NewReporteMantenimientoGeneradoRouter(engine)
	reporteMantenimientoGeneradoRoutes.Run()

	seguimientoFallaCriticaRoutes := seguimientoFallaCritica.NewSeguimientoFallaCriticaRouter(engine)
	seguimientoFallaCriticaRoutes.Run()

	tipoMantenimientoRoutes := tipoMantenimiento.NewTipoMantenimientoRouter(engine)
	tipoMantenimientoRoutes.Run()

	engine.Run(":8000")
}