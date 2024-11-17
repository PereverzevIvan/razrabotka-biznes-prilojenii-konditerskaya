package models

type Product struct {
	ID int `json:"id" gorm:"Column:id"`

	Name          string `json:"name" gorm:"Column:name"`
	IsSemiproduct bool   `json:"is_semiproduct" gorm:"Column:is_semiproduct"`
	Sizes         string `json:"sizes" gorm:"Column:sizes"`

	RecipeSemiProducts []*RecipeSemiproduct `json:"semiproducts"`
	RecipeComponents   []*RecipeComponents  `json:"recipe_components"`
}

func (Product) TableName() string {
	return "products"
}
