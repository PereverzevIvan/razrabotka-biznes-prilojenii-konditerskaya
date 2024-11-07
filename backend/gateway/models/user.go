package models

type User struct {
	ID       int    `json:"id" gorm:"Column:id"`
	RoleID   int    `json:"role_id" gorm:"Column:role_id"`
	Role     *Role  `json:"role" gorm:"foreignKey:RoleID"`
	Login    string `json:"login" gorm:"Column:login"`
	Password string `json:"-" gorm:"Column:password"`

	// OfficeID  int       `json:"office_id" gorm:"Column:OfficeID"`
	FirstName  string `json:"first_name" gorm:"Column:first_name"`
	LastName   string `json:"last_name" gorm:"Column:last_name"`
	Patronymic string `json:"patronymic" gorm:"Column:patronymic"`
	// Birthday  time.Time `json:"birthday" gorm:"Column:Birthdate"`
	// Active bool `json:"active" gorm:"Column:Active"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) GetID() int {
	return u.ID
}

func (u *User) GetRole() string {
	return ERole(u.RoleID).Name()
}
