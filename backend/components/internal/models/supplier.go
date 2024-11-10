package models

type Supplier struct {
	ID                  int     `json:"id" gorm:"Column:id"`
	Name                string  `json:"name" gorm:"Column:name"`
	Address             *string `json:"address" gorm:"Column:address"`
	DeliveryTimeMinutes int     `json:"delivery_time_minutes" gorm:"Column:delivery_time_minutes"`
}

func (Supplier) TableName() string {
	return "suppliers"
}
