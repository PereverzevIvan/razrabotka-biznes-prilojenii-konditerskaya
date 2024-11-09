package models

type ToolType struct {
	ID   int    `json:"id" gorm:"Column:id"`
	Name string `json:"name" gorm:"Column:name"`
}

func (t *ToolType) TableName() string {
	return "tool_types"
}
