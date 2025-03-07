package tool_failure_handler

import (
	tool_failure_deps "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool_failure/deps"
	tool_failure_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool_failure/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/storage"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool_failure"
	"google.golang.org/grpc"
)

type IToolFailureUsecase interface {
	Create(params *tool_failure_params.CreateParams) (*models.ToolFailure, error)
	AddFixedAt(params *tool_failure_params.AddFixedAtParams) error
	GetAll() ([]models.ToolFailure, error)
	GetAllReasons() ([]models.ToolFailureReason, error)
}

type ToolFailureHandler struct {
	tool_failure.UnimplementedToolFailureServiceServer
	toolFailureUsecase IToolFailureUsecase
}

func RegisterHandler(
	s grpc.ServiceRegistrar,
	storage *storage.Storage,
) {
	deps := tool_failure_deps.NewDeps(storage.DB, storage.JwtConfig)

	handler := &ToolFailureHandler{
		toolFailureUsecase: deps.ToolFailureUsecase(),
	}

	tool_failure.RegisterToolFailureServiceServer(s, handler)
}
