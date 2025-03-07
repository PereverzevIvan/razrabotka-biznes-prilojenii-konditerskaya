package tool_repo_grpc

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool"

type ToolRepo struct {
	client tool.ToolServiceClient
}

func NewToolRepo(client tool.ToolServiceClient) *ToolRepo {
	return &ToolRepo{client: client}
}
