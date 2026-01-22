package infrastructure

import (
	"github.com/jackc/pgx/v5/pgxpool"

	rolApp "github.com/vicpoo/API_recolecta/src/rol/application"
	rolController "github.com/vicpoo/API_recolecta/src/rol/infrastructure/controller"
	rolPostgres "github.com/vicpoo/API_recolecta/src/rol/infrastructure/postgres"
)

type RolDependencies struct {
	Create *rolController.RolController
	List   *rolController.RolController
	Update *rolController.RolController
	Delete *rolController.RolController
}

func NewRolDependencies(db *pgxpool.Pool) *rolController.RolController {
	repo := rolPostgres.NewRolRepository(db)

	createUC := rolApp.NewCreateRol(repo)
	listUC := rolApp.NewListRol(repo)
	updateUC := rolApp.NewUpdateRol(repo)
	deleteUC := rolApp.NewDeleteRol(repo)

	return rolController.NewRolController(createUC, listUC, updateUC, deleteUC)
}
