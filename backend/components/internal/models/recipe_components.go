package models

type RecipeComponents struct {
	ProductID int `json:"product_id" gorm:"Column:product_id"`
	// Product     *Product   `json:"product" gorm:"foreignKey:ProductID"`
	ComponentID int        `json:"component_id" gorm:"Column:component_id"`
	Component   *Component `json:"component" gorm:"foreignKey:ComponentID"`

	Count int `json:"count" gorm:"Column:count"`
}

func (RecipeComponents) TableName() string {
	return "recipe_components"
}
