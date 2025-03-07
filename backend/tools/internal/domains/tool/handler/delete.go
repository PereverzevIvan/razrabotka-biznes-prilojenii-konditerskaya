package tool_handler

import (
	"context"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool"
	"github.com/gofiber/fiber/v3/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *ToolHandler) Delete(ctx context.Context, req *tool.ToolDeleteRequest) (*tool.ToolDeleteResponse, error) {
	id := int(req.Id)
	if id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid tool id")
	}

	err := h.toolUsecase.Delete(id)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &tool.ToolDeleteResponse{}, nil
}
