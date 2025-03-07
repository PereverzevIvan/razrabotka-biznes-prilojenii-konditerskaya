package tool_failure_handler

import (
	"context"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool_failure"
	"github.com/gofiber/fiber/v3/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *ToolFailureHandler) GetAllReasons(ctx context.Context, req *tool_failure.ToolFailureGetAllReasonsRequest) (*tool_failure.ToolFailureGetAllReasonsResponse, error) {
	toolFailureReasons, err := h.toolFailureUsecase.GetAllReasons()
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	grpcToolFailureReasons := make([]*tool_failure.ToolFailureReason, len(toolFailureReasons))
	for i, toolFailureReason := range toolFailureReasons {
		grpcToolFailureReasons[i] = models.ToolFailureReasonToGRPC(&toolFailureReason)
	}

	return &tool_failure.ToolFailureGetAllReasonsResponse{
		ToolFailureReasons: grpcToolFailureReasons,
	}, nil
}
