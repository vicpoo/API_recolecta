//DeleteIncidenciaUseCase.go
package application

import (
	repositories "github.com/vicpoo/API_recolecta/src/incidencia/domain"
)

type DeleteIncidenciaUseCase struct {
	repo repositories.IIncidencia
}

func NewDeleteIncidenciaUseCase(repo repositories.IIncidencia) *DeleteIncidenciaUseCase {
	return &DeleteIncidenciaUseCase{repo: repo}
}

func (uc *DeleteIncidenciaUseCase) Run(id int32) error {
	// Primero obtenemos la incidencia
	incidencia, err := uc.repo.GetByID(id)
	if err != nil {
		return err
	}
	
	// Marcamos como eliminado (borrado lógico)
	incidencia.MarcarEliminado()
	
	// Actualizamos en la base de datos
	return uc.repo.Update(incidencia)
}