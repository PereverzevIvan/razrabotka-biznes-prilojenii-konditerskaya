package models

type ComponentType struct {
	ID int `json:"id" gorm:"Column:id"`

	ComponentCategoryID int                `json:"component_category_id" gorm:"Column:component_category_id"`
	ComponentCategory   *ComponentCategory `json:"component_category" gorm:"foreignKey:ComponentCategoryID"`

	Name string `json:"name" gorm:"Column:name"`
}

func (ComponentType) TableName() string {
	return "component_types"
}
