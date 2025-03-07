package tool_type_handler

import (
	"context"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool_type"

	"github.com/gofiber/fiber/v3/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *ToolTypeHandler) GetAll(ctx context.Context, req *tool_type.ToolTypeGetAllRequest) (*tool_type.ToolTypeGetAllResponse, error) {

	tool_types, err := h.toolTypeUsecase.GetAll()
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, "")
	}

	grpc_tool_types := make([]*tool_type.ToolType, len(tool_types))
	for i, tool_type := range tool_types {
		grpc_tool_types[i] = models.ToolTypeToGRPC(&tool_type)
	}

	return &tool_type.ToolTypeGetAllResponse{
		ToolType: grpc_tool_types,
	}, nil
}
