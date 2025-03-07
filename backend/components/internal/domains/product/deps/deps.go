package product_deps

import (
	component_repo_mysql "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/component/repo/mysql"
	product_repo_mysql "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/product/repo/mysql"
	product_usecase "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/product/usecase"
	purchased_component_repo_mysql "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/purchased_component/repo/mysql"
	supplier_component_repo_mysql "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/supplier_component/repo/mysql"
	tool_repo_grpc "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/tool/repo/grpc"
	tool_type_repo_grpc "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/tool_type/repo/grpc"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/proto/pkg/api/tool_type"
	"gorm.io/gorm"
)

type Deps struct {
	db *gorm.DB

	toolClient     tool.ToolServiceClient
	toolTypeClient tool_type.ToolTypeServiceClient

	componentRepo          *component_repo_mysql.ComponentRepo
	purchasedComponentRepo *purchased_component_repo_mysql.PurchasedComponentRepo
	supplierComponentRepo  *supplier_component_repo_mysql.SupplierComponentRepo
	toolTypeRepo           *tool_type_repo_grpc.ToolTypeRepo
	toolRepo               *tool_repo_grpc.ToolRepo

	productRepo    *product_repo_mysql.ProductRepo
	productUsecase *product_usecase.ProductUsecase
}

func NewDeps(
	db *gorm.DB,
	toolClient tool.ToolServiceClient,
	toolTypeClient tool_type.ToolTypeServiceClient,
) *Deps {
	return &Deps{
		db: db,

		toolClient:     toolClient,
		toolTypeClient: toolTypeClient,
	}
}

func (d *Deps) ComponentRepo() *component_repo_mysql.ComponentRepo {
	if d.componentRepo == nil {
		d.componentRepo = component_repo_mysql.NewComponentRepo(d.db)
	}
	return d.componentRepo
}
func (d *Deps) PurchasedComponentRepo() *purchased_component_repo_mysql.PurchasedComponentRepo {
	if d.purchasedComponentRepo == nil {
		d.purchasedComponentRepo = purchased_component_repo_mysql.NewPurchasedComponentRepo(d.db)
	}
	return d.purchasedComponentRepo
}
func (d *Deps) SupplierComponentRepo() *supplier_component_repo_mysql.SupplierComponentRepo {
	if d.supplierComponentRepo == nil {
		d.supplierComponentRepo = supplier_component_repo_mysql.NewSupplierComponentRepo(d.db)
	}
	return d.supplierComponentRepo
}
func (d *Deps) ToolTypeRepo() *tool_type_repo_grpc.ToolTypeRepo {
	if d.toolTypeRepo == nil {
		d.toolTypeRepo = tool_type_repo_grpc.NewToolTypeRepo(d.toolTypeClient)
	}
	return d.toolTypeRepo
}
func (d *Deps) ToolRepo() *tool_repo_grpc.ToolRepo {
	if d.toolRepo == nil {
		d.toolRepo = tool_repo_grpc.NewToolRepo(d.toolClient)
	}
	return d.toolRepo
}

func (d *Deps) ProductRepo() *product_repo_mysql.ProductRepo {
	if d.productRepo == nil {
		d.productRepo = product_repo_mysql.NewProductRepo(d.db)
	}
	return d.productRepo
}
func (d *Deps) ProductUsecase() *product_usecase.ProductUsecase {
	if d.productUsecase == nil {
		d.productUsecase = product_usecase.NewProductUsecase(
			d.ProductRepo(),
			d.ComponentRepo(),
			d.PurchasedComponentRepo(),
			d.SupplierComponentRepo(),
			d.ToolTypeRepo(),
			d.ToolRepo(),
		)
	}
	return d.productUsecase
}
