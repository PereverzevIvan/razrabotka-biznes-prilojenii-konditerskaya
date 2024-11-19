package models

type SupplierComponent struct {
	SupplierID  int     `json:"supplier_id" gorm:"Column:supplier_id"`
	ComponentID int     `json:"component_id" gorm:"Column:component_id"`
	Price       float64 `json:"price" gorm:"Column:price"`
}

func (SupplierComponent) TableName() string {
	return "supplier_components"
}
