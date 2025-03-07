package tool_failure_handler

import (
	"context"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool_failure"
	"github.com/gofiber/fiber/v3/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *ToolFailureHandler) GetAll(ctx context.Context, req *tool_failure.ToolFailureGetAllRequest) (*tool_failure.ToolFailureGetAllResponse, error) {
	toolFailures, err := h.toolFailureUsecase.GetAll()
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	grpcToolFailures := make([]*tool_failure.ToolFailure, len(toolFailures))
	for i, toolFailure := range toolFailures {
		grpcToolFailures[i] = models.ToolFailureToGRPC(&toolFailure)
	}

	return &tool_failure.ToolFailureGetAllResponse{
		ToolFailures: grpcToolFailures,
	}, nil
}
