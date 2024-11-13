package models

type ToolFailureReason struct {
	ID   int    `json:"id" gorm:"Column:id"`
	Name string `json:"name" gorm:"Column:name"`
}

func (ToolFailureReason) TableName() string {
	return "tool_failure_reasons"
}
