package tool_type_repo_grpc

import (
	"context"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool_type"
)

func (r *ToolTypeRepo) GetByID(ctx context.Context, id int) (*models.ToolType, error) {
	req := &tool_type.ToolTypeGetByIDRequest{
		Id: int32(id),
	}

	resp, err := r.client.GetByID(ctx, req)
	if err != nil {
		return nil, err
	}

	return models.ToolTypeFromGRPC(resp.ToolType), nil
}
