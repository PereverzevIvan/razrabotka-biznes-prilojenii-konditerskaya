package tool_handler

import (
	tool_deps "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool/deps"
	tool_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/domains/tool/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/internal/models"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/tools/storage"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool"
	"google.golang.org/grpc"
)

type IToolUsecase interface {
	GetAll(params *tool_params.GetAllParams) ([]models.Tool, error)
	GetByID(tool_id int) (*models.Tool, error)
	Add(params *tool_params.ToolAddParams) (*models.Tool, error)
	Edit(tool_id int, params *tool_params.ToolEditParams) (*models.Tool, error)
	Delete(tool_id int) error
}

type ToolHandler struct {
	tool.UnimplementedToolServiceServer
	toolUsecase IToolUsecase
}

func RegisterHandler(
	s grpc.ServiceRegistrar,
	storage *storage.Storage,
) {

	deps := tool_deps.NewDeps(storage.DB)

	handler := &ToolHandler{
		toolUsecase: deps.ToolUsecase(),
	}

	tool.RegisterToolServiceServer(s, handler)
}
