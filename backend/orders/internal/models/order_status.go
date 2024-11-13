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
	OrderStatusUndefined EOrderStatus = iota
	OrderStatusNew
	OrderStatusCanceled
	OrderStatusAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
	OrderStatusPurchase
	OrderStatusProduction
	OrderStatusControl
	OrderStatusReady
	OrderStatusCompleted
)
