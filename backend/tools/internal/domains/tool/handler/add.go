package tool_handler

import (
	"context"
	"fmt"

	tool_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool"
	"github.com/gofiber/fiber/v3/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *ToolHandler) Add(ctx context.Context, req *tool.ToolAddRequest) (*tool.ToolAddResponse, error) {

	params := tool_params.ToolAddParamsFromGRPC(req)
	if validate_errs := params.Validate(); len(validate_errs) != 0 {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprint(validate_errs))
	}

	toolModel, err := h.toolUsecase.Add(params)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &tool.ToolAddResponse{
		Tool: models.ToolToGRPC(toolModel),
	}, nil
}
