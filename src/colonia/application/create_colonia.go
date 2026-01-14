package application

import (
	"time"

	"github.com/vicpoo/API_recolecta/src/colonia/domain"
)

type CreateColonia struct {
	repo domain.ColoniaRepository
}

func NewCreateColonia(repo domain.ColoniaRepository) *CreateColonia {
	return &CreateColonia{repo}
}

func (uc *CreateColonia) Execute(c *domain.Colonia) error {
	c.CreatedAt = time.Now()
	return uc.repo.Create(c)
}
