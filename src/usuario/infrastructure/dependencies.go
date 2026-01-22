package infrastructure

import (
	"github.com/jackc/pgx/v5/pgxpool"

	usuarioApp "github.com/vicpoo/API_recolecta/src/usuario/application"
	usuarioController "github.com/vicpoo/API_recolecta/src/usuario/infrastructure/controller"
	usuarioPostgres "github.com/vicpoo/API_recolecta/src/usuario/infrastructure/postgres"
)

type UsuarioDependencies struct {
	Create *usuarioController.AddUsersController
	Delete *usuarioController.DeleteUsersController
	Get    *usuarioController.ViewOneUsersController
	List   *usuarioController.ViewAllUsersController
	Login  *usuarioController.LoginUsersController
}

func NewUsuarioDependencies(db *pgxpool.Pool) *UsuarioDependencies {
	repo := usuarioPostgres.NewUsuarioPostgresRepository(db)

	createUC := usuarioApp.NewSaveUser(repo)
	deleteUC := usuarioApp.NewDeleteUser(repo)
	getUC := usuarioApp.NewViewOneUser(repo)
	listUC := usuarioApp.NewViewAllUser(repo)
	loginUC := usuarioApp.NewLoginUser(repo)

	return &UsuarioDependencies{
		Create: usuarioController.NewAddUsersController(createUC),
		Delete: usuarioController.NewDeleteUsersController(deleteUC),
		Get:    usuarioController.NewViewOneUsersController(getUC),
		List:   usuarioController.NewViewAllUsersController(listUC),
		Login:  usuarioController.NewLoginUsersController(loginUC),
	}
}
