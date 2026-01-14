package application

import "github.com/vicpoo/API_recolecta/src/rol/domain"

type ListRoles struct {
	repo domain.RoleRepository
}

func NewListRoles(repo domain.RoleRepository) *ListRoles {
	return &ListRoles{repo}
}

func (uc *ListRoles) Execute() ([]domain.Role, error) {
	return uc.repo.GetAll()
}
