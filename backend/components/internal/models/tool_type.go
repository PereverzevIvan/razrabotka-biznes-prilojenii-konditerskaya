package models

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool_type"

type ToolType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ToolTypeFromGRPC(toolType *tool_type.ToolType) *ToolType {
	if toolType == nil {
		return nil
	}
	return &ToolType{
		ID:   int(toolType.Id),
		Name: toolType.Name,
	}
}
