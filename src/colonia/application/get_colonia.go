package application

import "github.com/vicpoo/API_recolecta/src/colonia/domain"

type GetColonia struct {
	repo domain.ColoniaRepository
}

func NewGetColonia(repo domain.ColoniaRepository) *GetColonia {
	return &GetColonia{repo}
}

func (uc *GetColonia) Execute(id int) (*domain.Colonia, error) {
	return uc.repo.GetByID(id)
}
