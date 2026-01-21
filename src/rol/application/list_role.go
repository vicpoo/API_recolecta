package application

import "github.com/vicpoo/API_recolecta/src/rol/domain"

type ListRol struct {
	repo domain.RolRepository
}

func NewListRol(r domain.RolRepository) *ListRol {
	return &ListRol{r}
}

func (uc *ListRol) Execute() ([]domain.Rol, error) {
	return uc.repo.List()
}
