package tool_type_repo_grpc

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool_type"
)

type ToolTypeRepo struct {
	client tool_type.ToolTypeServiceClient
}

func NewToolTypeRepo(client tool_type.ToolTypeServiceClient) *ToolTypeRepo {
	return &ToolTypeRepo{
		client: client,
	}
}
