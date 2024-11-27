package services_product

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/services"

type productService struct {
	productRepo            services.IProductRepo
	componentRepo          services.IComponentRepo
	purchasedComponentRepo services.IPurchasedComponentRepo
	supplierComponentRepo  services.ISupplierComponentRepo
	toolTypeRepo           services.IToolTypeRepo
	toolRepo               services.IToolRepo
}

func NewProductService(
	productRepo services.IProductRepo,
	componentRepo services.IComponentRepo,
	purchasesdComponentRepo services.IPurchasedComponentRepo,
	supplierComponentRepo services.ISupplierComponentRepo,
	toolTypeRepo services.IToolTypeRepo,
	toolRepo services.IToolRepo,
) *productService {
	service := &productService{
		productRepo:            productRepo,
		componentRepo:          componentRepo,
		purchasedComponentRepo: purchasesdComponentRepo,
		supplierComponentRepo:  supplierComponentRepo,
		toolTypeRepo:           toolTypeRepo,
		toolRepo:               toolRepo,
	}
	return service
}
