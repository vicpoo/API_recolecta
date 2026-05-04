package infrastructure

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/vicpoo/API_recolecta/src/empleado/application"
	"github.com/vicpoo/API_recolecta/src/empleado/infrastructure/controller"
	"github.com/vicpoo/API_recolecta/src/empleado/infrastructure/repository"
)

type EmpleadoDependencies struct {
	CreateEmpleadoController *controller.CreateEmpleadoController
	ListEmpleadoController   *controller.ListEmpleadoController
	GetEmpleadoController    *controller.GetEmpleadoController
	UpdateEmpleadoController *controller.UpdateEmpleadoController
	DeleteEmpleadoController *controller.DeleteEmpleadoController
	LoginEmpleadoController  *controller.LoginEmpleadoController
}

func InitEmpleadoDependencies(db *pgxpool.Pool) *EmpleadoDependencies {
	empleadoRepo := repository.NewEmpleadoPostgresRepository(db)
	createUseCase := application.NewCreateEmpleado(empleadoRepo)
	listUseCase := application.NewListEmpleado(empleadoRepo)
	getUseCase := application.NewGetEmpleado(empleadoRepo)
	updateUseCase := application.NewUpdateEmpleado(empleadoRepo)
	deleteUseCase := application.NewDeleteEmpleado(empleadoRepo)
	loginUseCase := application.NewLoginEmpleado(empleadoRepo)

	createController := controller.NewCreateEmpleadoController(createUseCase)
	listController := controller.NewListEmpleadoController(listUseCase)
	getController := controller.NewGetEmpleadoController(getUseCase)
	updateController := controller.NewUpdateEmpleadoController(updateUseCase)
	deleteController := controller.NewDeleteEmpleadoController(deleteUseCase)
	loginController := controller.NewLoginEmpleadoController(loginUseCase)

	return &EmpleadoDependencies{
		CreateEmpleadoController: createController,
		ListEmpleadoController:   listController,
		GetEmpleadoController:    getController,
		UpdateEmpleadoController: updateController,
		DeleteEmpleadoController: deleteController,
		LoginEmpleadoController:  loginController,
	}
}
