package tool_params

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool"

type GetAllParams struct {
	ToolTypeID      int  `json:"tool_type_id"`
	OnlyServiceable bool `json:"only_serviceable"`
}

func GetAllParamsToGRPC(params *GetAllParams) *tool.ToolGetAllRequest {
	req := &tool.ToolGetAllRequest{
		OnlyServiceable: &params.OnlyServiceable,
	}

	toolType := int32(params.ToolTypeID)
	req.ToolType = &toolType

	return req
}
