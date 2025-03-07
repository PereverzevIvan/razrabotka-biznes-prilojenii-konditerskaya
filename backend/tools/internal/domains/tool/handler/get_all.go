package tool_handler

import (
	"context"

	tool_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool"
	"github.com/gofiber/fiber/v3/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *ToolHandler) GetAll(ctx context.Context, req *tool.ToolGetAllRequest) (*tool.ToolGetAllResponse, error) {

	params := tool_params.GetAllParamsFromGRPC(req)

	tools, err := h.toolUsecase.GetAll(params)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	toolsGRPC := make([]*tool.Tool, len(tools))
	for i, tool := range tools {
		toolsGRPC[i] = models.ToolToGRPC(&tool)
	}

	return &tool.ToolGetAllResponse{
		Tools: toolsGRPC,
	}, nil
}
