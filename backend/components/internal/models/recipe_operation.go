package models

type RecipeOperation struct {
	ID int `json:"id" gorm:"Column:id"`

	ProductID int      `json:"product_id" gorm:"Column:product_id"`
	Product   *Product `json:"product" gorm:"foreignKey:ProductID"`

	ToolTypeID int       `json:"tool_type_id" gorm:"Column:tool_type_id"`
	ToolType   *ToolType `json:"tool_type" gorm:"foreignKey:ToolTypeID"`

	Desctiption     string `json:"description" gorm:"Column:description"`
	DurationMinutes int    `json:"duration_minutes" gorm:"Column:duration_minutes"`
	OrderIdx        int    `json:"order_idx" gorm:"Column:order_idx"`
}

func (RecipeOperation) TableName() string {
	return "recipe_operations"
}
