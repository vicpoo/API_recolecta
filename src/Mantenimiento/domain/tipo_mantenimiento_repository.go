//tipo_mantenimiento_repository.go
package domain

import(
	"github.com/vicpoo/API_recolecta/src/Mantenimiento/domain/entities"
)

type ITipoMantenimiento interface{
	Save(tipoMantenimiento *entities.TipoMantenimiento) error
	Update(tipoMantenimiento *entities.TipoMantenimiento) error
	Delete(id int32) error
	GetAll()([]entities.TipoMantenimiento,error)
	GetByID(id int32)(*entities.TipoMantenimiento,error)
}