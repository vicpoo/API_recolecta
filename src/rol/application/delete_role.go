package application

import "github.com/vicpoo/API_recolecta/src/rol/domain"

type DeleteRol struct {
	repo domain.RolRepository
}

func NewDeleteRol(r domain.RolRepository) *DeleteRol {
	return &DeleteRol{r}
}

func (uc *DeleteRol) Execute(id int) error {
	return uc.repo.Delete(id)
}
