package models

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool_type"

type ToolType struct {
	ID   int    `json:"id" gorm:"Column:id"`
	Name string `json:"name" gorm:"Column:name"`
}

func (t *ToolType) TableName() string {
	return "tool_types"
}

func ToolTypeToGRPC(toolType *ToolType) *tool_type.ToolType {
	if toolType == nil {
		return nil
	}
	return &tool_type.ToolType{
		Id:   int32(toolType.ID),
		Name: toolType.Name,
	}
}
