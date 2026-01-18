package domain

type RoleRepository interface {
	Create(role *Role) error
	GetByID(roleID int) (*Role, error)
	GetAll() ([]*Role, error)
	Update(role *Role) error
}