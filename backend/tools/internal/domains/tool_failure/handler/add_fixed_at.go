package tool_failure_handler

import (
	"context"
	"fmt"

	tool_failure_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool_failure/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool_failure"
	"github.com/gofiber/fiber/v3/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *ToolFailureHandler) AddFixedAt(ctx context.Context, req *tool_failure.ToolFailureAddFixedAtRequest) (*tool_failure.ToolFailureAddFixedAtResponse, error) {
	params := tool_failure_params.AddFixedAddParamsFromGRPC(req)
	if validate_errs := params.Validate(); len(validate_errs) != 0 {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprint(validate_errs))
	}

	err := h.toolFailureUsecase.AddFixedAt(params)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &tool_failure.ToolFailureAddFixedAtResponse{}, nil
}
