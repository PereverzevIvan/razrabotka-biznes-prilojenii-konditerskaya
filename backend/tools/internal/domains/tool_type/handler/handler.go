package tool_type_handler

import (
	tool_type_deps "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool_type/deps"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/storage"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool_type"

	"google.golang.org/grpc"
)

type IToolTypeUsecase interface {
	GetAll() ([]models.ToolType, error)
	GetByID(id int) (*models.ToolType, error)
}

type ToolTypeHandler struct {
	tool_type.UnimplementedToolTypeServiceServer

	toolTypeUsecase IToolTypeUsecase
}

func RegisterHandler(
	s grpc.ServiceRegistrar,
	storage *storage.Storage,
) {
	deps := tool_type_deps.NewDeps(storage.DB)

	handler := &ToolTypeHandler{
		toolTypeUsecase: deps.ToolTypeUsecase(),
	}

	tool_type.RegisterToolTypeServiceServer(s, handler)
}
