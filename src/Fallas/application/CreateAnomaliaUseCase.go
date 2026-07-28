// CreateAnomaliaUseCase.go
package application

import (
	"time"

	repositories "github.com/vicpoo/API_recolecta/src/Fallas/domain"
	"github.com/vicpoo/API_recolecta/src/Fallas/domain/entities"
	alertaDomain "github.com/vicpoo/API_recolecta/src/alerta_usuario/domain"
)

type CreateAnomaliaUseCase struct {
	repo       repositories.IAnomalia
	alertaRepo alertaDomain.AlertaUsuarioRepository
}

func NewCreateAnomaliaUseCase(repo repositories.IAnomalia, alertaRepo alertaDomain.AlertaUsuarioRepository) *CreateAnomaliaUseCase {
	return &CreateAnomaliaUseCase{repo: repo, alertaRepo: alertaRepo}
}

func (uc *CreateAnomaliaUseCase) Run(anomalia *entities.Anomalia) (*entities.Anomalia, error) {
	err := uc.repo.Save(anomalia)
	if err != nil {
		return nil, err
	}

	// Si es una anomalía crítica o incidencia, alertar a supervisores
	tipo := string(anomalia.TipoAnomalia)
	if tipo == "REPORTE_FALLA_CRITICA" || tipo == "INCIDENCIA" {
		creadoPor := 0
		if anomalia.ConductorID != nil {
			creadoPor = int(*anomalia.ConductorID)
		}
		alerta := &alertaDomain.AlertaUsuario{
			Titulo:    "Alerta Operativa: " + tipo,
			Mensaje:   anomalia.Descripcion,
			UsuarioID: 1, // Administrador general
			Leida:     false,
			CreadoPor: creadoPor,
			CreatedAt: time.Now(),
		}
		_ = uc.alertaRepo.Create(alerta)
	}

	return anomalia, nil
}