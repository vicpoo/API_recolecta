package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	camionUseCases "github.com/vicpoo/API_recolecta/src/Camion/application"
	camionAdapters "github.com/vicpoo/API_recolecta/src/Camion/infraestructure/adapters"
	camionControllers "github.com/vicpoo/API_recolecta/src/Camion/infraestructure/controllers"
	camionRoutes "github.com/vicpoo/API_recolecta/src/Camion/infraestructure/routes"
	estadoCamionUseCases "github.com/vicpoo/API_recolecta/src/EstadoCamion/application"
	estadoCamionAdapters "github.com/vicpoo/API_recolecta/src/EstadoCamion/infraestructure/adapters"
	estadoCamionControllers "github.com/vicpoo/API_recolecta/src/EstadoCamion/infraestructure/controllers"
	estadoCamionRoutes "github.com/vicpoo/API_recolecta/src/EstadoCamion/infraestructure/routes"
	tipoCamionUseCases "github.com/vicpoo/API_recolecta/src/TipoCamion/application"
	tipoCamionAdapters "github.com/vicpoo/API_recolecta/src/TipoCamion/infraestructure/adapters"
	tipoCamionControllers "github.com/vicpoo/API_recolecta/src/TipoCamion/infraestructure/controllers"
	tipoCamionRoutes "github.com/vicpoo/API_recolecta/src/TipoCamion/infraestructure/routes"
	"github.com/vicpoo/API_recolecta/src/core"
)

//archivo para hacer las instancias de los controllers, casos de uso y repositories, etc.
func InitDependencies() {
	if errEnv := godotenv.Load(); errEnv != nil {
		log.Fatal("error al cargar el .env")
	}

	engine := gin.Default()
	engine.Use(core.CORSMiddleware())


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

	
	engine.Run(":8080")
}