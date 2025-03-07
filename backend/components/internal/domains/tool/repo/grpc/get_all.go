package tool_repo_grpc

import (
	"context"

	tool_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/tool/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
)

func (r *ToolRepo) GetAll(ctx context.Context, params *tool_params.GetAllParams) ([]models.Tool, error) {

	req := tool_params.GetAllParamsToGRPC(params)

	resp, err := r.client.GetAll(ctx, req)
	if err != nil {
		return nil, err
	}

	var tools []models.Tool
	for _, tool := range resp.Tools {
		tools = append(tools, *models.ToolFromGRPC(tool))
	}

	return tools, nil
}
