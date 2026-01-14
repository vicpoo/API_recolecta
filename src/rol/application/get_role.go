package application

import "github.com/vicpoo/API_recolecta/src/rol/domain"

type GetRole struct{
	repo domain.RoleRepository
}

func NewGetRole(repo domain.RoleRepository) *GetRole {
	return &GetRole{repo}
}

func (uc *GetRole) Execute(id int) (*domain.Role, error) {
	return uc.repo.GetByID(id)
}