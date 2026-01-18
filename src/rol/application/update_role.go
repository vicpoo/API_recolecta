package application

import "github.com/vicpoo/API_recolecta/src/rol/domain"

type UpdateRole struct {
	repo domain.RoleRepository
}

func NewUpdateRole(repo domain.RoleRepository) *UpdateRole {
	return &UpdateRole{repo}
}

func (uc *UpdateRole) Execute(id int, nombre string) error {
	role := &domain.Role{
		RoleID: id,
		Nombre: nombre,
	}
	return uc.repo.Update(role)
}
