package models

type Component struct {
	ID int `json:"id" gorm:"Column:id"`

	ComponentTypeID int            `json:"type_id" gorm:"Column:component_type_id"`
	ComponentType   *ComponentType `json:"type" gorm:"foreignKey:ComponentTypeID"`

	Name    string `json:"name" gorm:"Column:name"`
	Article string `json:"article" gorm:"Column:article"`

	MeasureUnit string `json:"measure_unit" gorm:"Column:measure_unit"`
}

func (Component) TableName() string {
	return "components"
}
