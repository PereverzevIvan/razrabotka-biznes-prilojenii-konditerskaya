package tool_failure_handler

import (
	"context"
	"fmt"

	tool_failure_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool_failure/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool_failure"
	"github.com/gofiber/fiber/v3/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *ToolFailureHandler) Create(ctx context.Context, req *tool_failure.ToolFailureCreateRequest) (*tool_failure.ToolFailureCreateResponse, error) {

	master_id, ok := ctx.Value("user_id").(int)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "user_id not found in context")
	}

	params := tool_failure_params.CreateParamsFromGRPC(req)
	params.MasterID = master_id

	if validate_errs := params.Validate(); len(validate_errs) != 0 {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprint(validate_errs))
	}

	toolFailure, err := h.toolFailureUsecase.Create(params)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &tool_failure.ToolFailureCreateResponse{
		ToolFailure: models.ToolFailureToGRPC(toolFailure),
	}, nil
}
