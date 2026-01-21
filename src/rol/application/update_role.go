package application

import "github.com/vicpoo/API_recolecta/src/rol/domain"

type UpdateRol struct {
	repo domain.RolRepository
}

func NewUpdateRol(r domain.RolRepository) *UpdateRol {
	return &UpdateRol{r}
}

func (uc *UpdateRol) Execute(id int, nombre string) error {
	return uc.repo.Update(id, nombre)
}
