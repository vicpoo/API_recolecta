package application

import "github.com/vicpoo/API_recolecta/src/colonia/domain"

type DeleteColonia struct {
	repo domain.ColoniaRepository
}

func NewDeleteColonia(repo domain.ColoniaRepository) *DeleteColonia {
	return &DeleteColonia{repo}
}

func (uc *DeleteColonia) Execute(id int) error {
	return uc.repo.Delete(id)
}
