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
	KRoleUser
	KRoleAdmin
)

func (e ERole) Name() string {
	switch e {
	case KRoleUser:
		return "user"
	case KRoleAdmin:
		return "admin"
	default:
		return "none"
	}
}
