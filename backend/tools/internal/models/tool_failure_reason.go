package models

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool_failure"

type ToolFailureReason struct {
	ID   int    `json:"id" gorm:"Column:id"`
	Name string `json:"name" gorm:"Column:name"`
}

func (ToolFailureReason) TableName() string {
	return "tool_failure_reasons"
}

func GRPCToolFailureReason(toolFailureReason *ToolFailureReason) *tool_failure.ToolFailureReason {
	return &tool_failure.ToolFailureReason{
		Id:   int32(toolFailureReason.ID),
		Name: toolFailureReason.Name,
	}
}

func ToolFailureReasonToGRPC(toolFailureReason *ToolFailureReason) *tool_failure.ToolFailureReason {
	return &tool_failure.ToolFailureReason{
		Id:   int32(toolFailureReason.ID),
		Name: toolFailureReason.Name,
	}
}
