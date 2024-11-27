package params_tool

type GetAllParams struct {
	ToolTypeID      int  `json:"tool_type_id"`
	OnlyServiceable bool `json:"only_serviceable"`
}
