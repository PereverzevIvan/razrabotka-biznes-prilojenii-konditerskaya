package models

type RecipeSemiproduct struct {
	ProductID int `json:"product_id" gorm:"Column:product_id"`
	// Product       *Product `json:"product" gorm:"foreignKey:ProductID"`
	SemiproductID int      `json:"semiproduct_id" gorm:"Column:semiproduct_id"`
	Semiproduct   *Product `json:"semiproduct" gorm:"foreignKey:SemiproductID"`
	Count         int      `json:"count" gorm:"Column:count"`
}

func (RecipeSemiproduct) TableName() string {
	return "recipe_semiproducts"
}
