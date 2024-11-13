package models

import "time"

type Order struct {
	ID int `json:"id" gorm:"Column:id"`

	ProductID int `json:"product_id" gorm:"Column:product_id"`
	// Product   *Product

	CustomerID int   `json:"customer_id" gorm:"Column:customer_id"`
	Customer   *User `json:"customer" gorm:"foreignKey:CustomerID"`

	ManagerID *int  `json:"manager_id" gorm:"Column:manager_id"`
	Manager   *User `json:"manager" gorm:"foreignKey:ManagerID"`

	StatusID int          `json:"status_id" gorm:"Column:status_id"`
	Status   *OrderStatus `json:"status" gorm:"foreignKey:StatusID"`

	Number string   `json:"number" gorm:"Column:number"`
	Name   string   `json:"name" gorm:"Column:name"`
	Cost   *float64 `json:"cost" gorm:"Column:cost"`

	CreatedAt time.Time `json:"created_at" gorm:"Column:created_at"`
	// UpdatedAt           time.Time  `json:"updated_at" gorm:"Column:updated_at"`
	PlannedCompletionAt *time.Time `json:"planned_completion_at" gorm:"Column:planned_completion_at"`

	// ExampleImages []string `json:"example_images" gorm:"Column:example_images"`
}
