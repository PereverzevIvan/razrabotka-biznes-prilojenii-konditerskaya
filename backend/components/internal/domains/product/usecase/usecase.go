package product_usecase

import (
	"context"
	"mime/multipart"

	purchased_component_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/purchased_component/params"
	tool_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/domains/tool/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	"github.com/gofiber/fiber/v3"
)

type IProductRepo interface {
	GetAll() ([]models.Product, error)
	Create(product *models.Product) (*models.Product, error)
	Update(product *models.Product) error
	GetByID(id int) (*models.Product, error)
	SaveImage(ctx fiber.Ctx, id int, image *multipart.FileHeader) (string, error)
}

type IComponentRepo interface {
	// GetAll(component_category_id int) ([]models.Component, error)
	GetByID(id int) (*models.Component, error)
}

type IPurchasedComponentRepo interface {
	GetAll(component_category_id int, params *purchased_component_params.GetAllParams) ([]models.PurchasedComponent, error)
	GetAllTotalRows(component_category_id int, params *purchased_component_params.GetAllParams) (int64, error)
	GetAllTotalCount(component_category_id int, params *purchased_component_params.GetAllParams) (int64, error)
	GetAllTotalPrice(component_category_id int, params *purchased_component_params.GetAllParams) (float64, error)

	DeleteEmpty(purchased_component_id int) error

	Create(purchased_component *models.PurchasedComponent) error

	Edit(purchased_component *models.PurchasedComponent) error

	CountRemainingComponents(component_id int) (int64, error)
	DeductComponents(component_id int, deduct_count int) error
	CalcPriceOfRequiredCount(component_id int, count int) (float64, error)
}
type ISupplierComponentRepo interface {
	FastestDeliveryComponentPrice(component_id int) (float64, error)
	FastestDeliveryComponentTime(component_id int) (int, error)
}

type IToolTypeRepo interface {
	GetByID(ctx context.Context, id int) (*models.ToolType, error)
}

type IToolRepo interface {
	GetAll(ctx context.Context, params *tool_params.GetAllParams) ([]models.Tool, error)
}

type ProductUsecase struct {
	productRepo            IProductRepo
	componentRepo          IComponentRepo
	purchasedComponentRepo IPurchasedComponentRepo
	supplierComponentRepo  ISupplierComponentRepo
	toolTypeRepo           IToolTypeRepo
	toolRepo               IToolRepo
}

func NewProductUsecase(
	productRepo IProductRepo,
	componentRepo IComponentRepo,
	purchasesdComponentRepo IPurchasedComponentRepo,
	supplierComponentRepo ISupplierComponentRepo,
	toolTypeRepo IToolTypeRepo,
	toolRepo IToolRepo,
) *ProductUsecase {
	usecase := &ProductUsecase{
		productRepo:            productRepo,
		componentRepo:          componentRepo,
		purchasedComponentRepo: purchasesdComponentRepo,
		supplierComponentRepo:  supplierComponentRepo,
		toolTypeRepo:           toolTypeRepo,
		toolRepo:               toolRepo,
	}

	return usecase
}
