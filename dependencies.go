package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vicpoo/API_recolecta/src/core"
	camionUseCases "github.com/vicpoo/API_recolecta/src/Camion/application"
	camionAdapters "github.com/vicpoo/API_recolecta/src/Camion/infraestructure/adapters"
	camionControllers "github.com/vicpoo/API_recolecta/src/Camion/infraestructure/controllers"
	camionRoutes "github.com/vicpoo/API_recolecta/src/Camion/infraestructure/routes"
	estadoCamionUseCases "github.com/vicpoo/API_recolecta/src/EstadoCamion/application"
	estadoCamionAdapters "github.com/vicpoo/API_recolecta/src/EstadoCamion/infraestructure/adapters"
	estadoCamionControllers "github.com/vicpoo/API_recolecta/src/EstadoCamion/infraestructure/controllers"
	estadoCamionRoutes "github.com/vicpoo/API_recolecta/src/EstadoCamion/infraestructure/routes"
	historialUseCases "github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/application"
	historialAdapters "github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/infraestructure/adapters"
	historialControllers "github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/infraestructure/controllers"
	historialRoutes "github.com/vicpoo/API_recolecta/src/HistorialAsignacionCamion/infraestructure/routes"
	puntoUseCases "github.com/vicpoo/API_recolecta/src/PuntoRecoleccion/application"
	puntoAdapters "github.com/vicpoo/API_recolecta/src/PuntoRecoleccion/infraestructure/adapters"
	puntoControllers "github.com/vicpoo/API_recolecta/src/PuntoRecoleccion/infraestructure/controllers"
	puntoRoutes "github.com/vicpoo/API_recolecta/src/PuntoRecoleccion/infraestructure/routes"
	rsApplication "github.com/vicpoo/API_recolecta/src/RellenoSanitario/application"
	rsAdapters "github.com/vicpoo/API_recolecta/src/RellenoSanitario/infraestructure/adapters"
	rsControllers "github.com/vicpoo/API_recolecta/src/RellenoSanitario/infraestructure/controllers"
	rsRoutes "github.com/vicpoo/API_recolecta/src/RellenoSanitario/infraestructure/routes"
	rutaUseCases "github.com/vicpoo/API_recolecta/src/Ruta/application"
	rutaAdapters "github.com/vicpoo/API_recolecta/src/Ruta/infraestructure/adapters"
	rutaControllers "github.com/vicpoo/API_recolecta/src/Ruta/infraestructure/controllers"
	rutaRoutes "github.com/vicpoo/API_recolecta/src/Ruta/infraestructure/routes"
	rutaCamionApp "github.com/vicpoo/API_recolecta/src/RutaCamion/application"
	rutaCamionAdapters "github.com/vicpoo/API_recolecta/src/RutaCamion/infraestructure/adapters"
	rutaCamionControllers "github.com/vicpoo/API_recolecta/src/RutaCamion/infraestructure/controllers"
	rutaCamionRoutes "github.com/vicpoo/API_recolecta/src/RutaCamion/infraestructure/routes"
	tipoCamionUseCases "github.com/vicpoo/API_recolecta/src/TipoCamion/application"
	tipoCamionAdapters "github.com/vicpoo/API_recolecta/src/TipoCamion/infraestructure/adapters"
	tipoCamionControllers "github.com/vicpoo/API_recolecta/src/TipoCamion/infraestructure/controllers"
	tipoCamionRoutes "github.com/vicpoo/API_recolecta/src/TipoCamion/infraestructure/routes"
    registroVaciadoAdapters "github.com/vicpoo/API_recolecta/src/RegistroVaciado/infraestructure/adapters"
    registroVaciadoApplication "github.com/vicpoo/API_recolecta/src/RegistroVaciado/application"
    registroVaciadoControllers "github.com/vicpoo/API_recolecta/src/RegistroVaciado/infraestructure/controllers"
    registroVaciadoRoutesPkg "github.com/vicpoo/API_recolecta/src/RegistroVaciado/infraestructure/routes"

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

//archivo para hacer las instancias de los controllers, casos de uso y repositories, etc.
func InitDependencies() {
	if errEnv := godotenv.Load(); errEnv != nil {
		log.Fatal("error al cargar el .env")
	}

	engine := gin.Default()
	engine.Use(core.CORSMiddleware())

	db := core.GetBD()

	//tipo camion
	tipoCamionRepository := tipoCamionAdapters.NewPosgres()
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


	//camion
	camionRepository := camionAdapters.NewPostgres()
	saveCamionUc :=  camionUseCases.NewSaveCamionUseCase(camionRepository)
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

	//estado camion
	estadoCamionRepository := estadoCamionAdapters.NewPostgres()

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
	historialRepository := historialAdapters.NewPostgres()

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

	rutaRepository := rutaAdapters.NewPostgres()

	createRutaUc := rutaUseCases.NewCreateRutaUseCase(rutaRepository)
	getAllRutaUc := rutaUseCases.NewListAllRutaUseCase(rutaRepository)
	getRutaByIdUc := rutaUseCases.NewGetRutaByIdUseCase(rutaRepository)
	updateRutaUc := rutaUseCases.NewUpdateRutaUseCase(rutaRepository)
	deleteRutaUc := rutaUseCases.NewDeleteRutaUseCase(rutaRepository)

	createRutaCtr := rutaControllers.NewCreateRutaController(createRutaUc)
	getAllRutaCtr := rutaControllers.NewGetAllRutaController(getAllRutaUc)
	getRutaByIdCtr := rutaControllers.NewGetRutaByIdController(getRutaByIdUc)
	updateRutaCtr := rutaControllers.NewUpdateRutaController(updateRutaUc)
	deleteRutaCtr := rutaControllers.NewDeleteRutaController(deleteRutaUc)

	rutaRoutes := rutaRoutes.NewRutaRoutes(
		engine,
		createRutaCtr,
		getAllRutaCtr,
		getRutaByIdCtr,
		updateRutaCtr,
		deleteRutaCtr,
	)

	rutaRoutes.Run()

    puntoRepository := puntoAdapters.NewPostgres()

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


	rellenoRepo := rsAdapters.NewPostgres()

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

repository := rutaCamionAdapters.NewPostgres()

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
registroVaciadoRepository := registroVaciadoAdapters.NewPostgres()

// ===============================
// USE CASES
// ===============================
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

	engine.Run(":8080")
}