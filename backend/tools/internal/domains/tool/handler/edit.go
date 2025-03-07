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

func (h *ToolHandler) Edit(ctx context.Context, req *tool.ToolEditRequest) (*tool.ToolEditResponse, error) {

	tool_id := int(req.GetId())
	if tool_id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid tool id")
	}

	params := tool_params.ToolEditParamsFromGRPC(req)

	toolModel, err := h.toolUsecase.Edit(tool_id, params)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	// return ctx.Status(fiber.StatusCreated).JSON(tool)
	return &tool.ToolEditResponse{
		Tool: models.ToolToGRPC(toolModel),
	}, nil
}
