package storage

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/configs"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool_type"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Storage struct {
	DB        *gorm.DB
	JwtConfig *configs.JWTConfig

	ToolClient     tool.ToolServiceClient
	ToolTypeClient tool_type.ToolTypeServiceClient
}

func NewStorage(
	dbConfig *configs.DBConfig,
	jwtConfig *configs.JWTConfig,
	toolConn *grpc.ClientConn,
) (*Storage, error) {
	var err error

	storage := &Storage{
		JwtConfig:      jwtConfig,
		ToolClient:     tool.NewToolServiceClient(toolConn),
		ToolTypeClient: tool_type.NewToolTypeServiceClient(toolConn),
	}

	storage.DB, err = NewDB(dbConfig)
	if err != nil {
		return nil, err
	}

	return storage, nil
}
