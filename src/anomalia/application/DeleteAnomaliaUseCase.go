// DeleteAnomaliaUseCase.go
package application

import repositories "github.com/vicpoo/API_recolecta/src/anomalia/domain"

type DeleteAnomaliaUseCase struct {
	repo repositories.IAnomalia
}

func NewDeleteAnomaliaUseCase(repo repositories.IAnomalia) *DeleteAnomaliaUseCase {
	return &DeleteAnomaliaUseCase{repo: repo}
}

func (uc *DeleteAnomaliaUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}