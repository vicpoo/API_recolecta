package application

import "github.com/vicpoo/API_recolecta/src/colonia/domain"

type UpdateColonia struct {
	repo domain.ColoniaRepository
}

func NewUpdateColonia(repo domain.ColoniaRepository) *UpdateColonia {
	return &UpdateColonia{repo}
}

func (uc *UpdateColonia) Execute(c *domain.Colonia) error {
	return uc.repo.Update(c)
}
