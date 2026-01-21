package application

import "github.com/vicpoo/API_recolecta/src/rol/domain"

type CreateRol struct {
	repo domain.RolRepository
}

func NewCreateRol(r domain.RolRepository) *CreateRol {
	return &CreateRol{r}
}

func (uc *CreateRol) Execute(nombre string) error {
	return uc.repo.Create(nombre)
}
