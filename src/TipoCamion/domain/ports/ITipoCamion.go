package ports

import "github.com/vicpoo/API_recolecta/src/TipoCamion/domain/entities"

type ITipoCamion interface {
	Save(tipoCamion *entities.TipoCamion) (*entities.TipoCamion, error)
	ListAll() ([]entities.TipoCamion, error)
	GetByName(nombre string) (*entities.TipoCamion, error) 
	Delete(id int32) error
}
