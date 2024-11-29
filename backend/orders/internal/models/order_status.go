package models

type OrderStatus struct {
	ID   int    `json:"id" gorm:"Column:id"`
	Name string `json:"name" gorm:"Column:name"`
}

func (o *OrderStatus) TableName() string {
	return "order_statuses"
}

type EOrderStatus int

const (
	KOrderStatusUndefined EOrderStatus = iota
	KOrderStatusNew
	KOrderStatusCanceled
	KOrderStatusCompilingRecipe
	KOrderStatusConfirmation
	KOrderStatusPurchase
	KOrderStatusProduction
	KOrderStatusControl
	KOrderStatusReady
	KOrderStatusCompleted
)
