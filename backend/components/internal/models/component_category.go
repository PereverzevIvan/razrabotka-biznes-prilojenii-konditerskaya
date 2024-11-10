package models

type ComponentCategory struct {
	ID   int    `json:"id" gorm:"Column:id"`
	Name string `json:"name" gorm:"Column:name"`
}

func (ComponentCategory) TableName() string {
	return "component_categories"
}

type EComponentCategory int

const (
	KComponentCategoryNone EComponentCategory = iota
	KComponentCategoryIngredients
	KComponentCategoryDecorations
)

func (e EComponentCategory) Name() string {
	switch e {
	case KComponentCategoryNone:
		return ""
	case KComponentCategoryIngredients:
		return "ingredients"
	case KComponentCategoryDecorations:
		return "decorations"
	default:
		return ""
	}
}
