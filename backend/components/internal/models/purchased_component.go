package models

import "time"

type PurchasedComponent struct {
	ID int `json:"id" gorm:"Column:id"`

	ComponentID int        `json:"component_id" gorm:"Column:component_id"`
	Component   *Component `json:"component" gorm:"foreignKey:ComponentID"`

	Quantity int `json:"quantity" gorm:"Column:quantity"`

	PurchasePrice float64 `json:"purchase_price" gorm:"Column:purchase_price"`

	ShelfLife *time.Time `json:"shelf_life" gorm:"Column:shelf_life"`
}

func (PurchasedComponent) TableName() string {
	return "purchased_components"
}
