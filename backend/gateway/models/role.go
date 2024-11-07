package models

type Role struct {
	ID   int    `json:"id" gorm:"Column:id"`
	Name string `json:"name" gorm:"Column:name"`
}

func (r *Role) TableName() string {
	return "roles"
}

type ERole int

const (
	KRoleNone ERole = iota
	KRoleSupplier
	KRoleClientManager
	KRolePurchaseManager
	KRoleMaster
	KRoleDirector
)

func (e ERole) Name() string {
	switch e {
	case KRoleSupplier:
		return "Заказчик"
	case KRoleClientManager:
		return "Менеджер по работе с клиентами"
	case KRolePurchaseManager:
		return "Менеджер по закупкам"
	case KRoleMaster:
		return "Мастер"
	case KRoleDirector:
		return "Директор"
	default:
		return "none"
	}
}
