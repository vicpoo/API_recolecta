package application

import "github.com/vicpoo/API_recolecta/src/colonia/domain"

type ListColonias struct {
	repo domain.ColoniaRepository
}

func NewListColonias(repo domain.ColoniaRepository) *ListColonias {
	return &ListColonias{repo}
}

func (uc *ListColonias) Execute() ([]domain.Colonia, error) {
	return uc.repo.GetAll()
}
