package tool_type_handler

import (
	"context"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool_type"
	"github.com/gofiber/fiber/v3/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *ToolTypeHandler) GetByID(ctx context.Context, req *tool_type.ToolTypeGetByIDRequest) (*tool_type.ToolTypeGetByIDResponse, error) {

	toolType, err := h.toolTypeUsecase.GetByID(int(req.Id))
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, "")
	}
	if toolType == nil {
		return nil, status.Error(codes.NotFound, "")
	}

	return &tool_type.ToolTypeGetByIDResponse{
		ToolType: models.ToolTypeToGRPC(toolType),
	}, nil
}
