package models

type QualityControl struct {
	ID int `json:"id" gorm:"Column:id"`

	OrderID int    `json:"order_id" gorm:"Column:order_id"`
	Order   *Order `json:"order" gorm:"foreignKey:OrderID"`

	MasterID int   `json:"master_id" gorm:"Column:master_id"`
	Master   *User `json:"master" gorm:"foreignKey:MasterID"`

	IsSatisfies bool `json:"is_satisfies" gorm:"Column:is_satisfies"`

	Comment   *string `json:"comment" gorm:"Column:comment"`
	Parameter *string `json:"parameter" gorm:"Column:parameter"`
}

func (*QualityControl) TableName() string {
	return "quality_controls"
}
