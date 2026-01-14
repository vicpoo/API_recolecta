package application

import "github.com/vicpoo/API_recolecta/src/rol/domain"

type CreateRole struct{
	repo domain.RoleRepository
}

func NewcreateRole(repo domain.RoleRepository) *CreateRole {
	return &CreateRole{repo}
}

func (uc *CreateRole) Execute(nombre string) error {
	role := &domain.Role{
		Nombre: nombre,
		Eliminado: false,
	}
	return uc.repo.Create(role)
}