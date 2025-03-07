package tool_handler

import (
	"context"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool"
	"github.com/gofiber/fiber/v3/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *ToolHandler) GetByID(ctx context.Context, req *tool.ToolGetByIDRequest) (*tool.ToolGetByIDResponse, error) {

	id := int(req.GetId())
	if id <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid tool id")
	}

	toolModel, err := h.toolUsecase.GetByID(id)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if toolModel == nil {
		return nil, status.Error(codes.NotFound, "tool not found")
	}

	return &tool.ToolGetByIDResponse{
		Tool: models.ToolToGRPC(toolModel),
	}, nil
}
