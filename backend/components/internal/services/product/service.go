package services_product

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/services"

type productService struct {
	productRepo   services.IProductRepo
	componentRepo services.IComponentRepo
	toolTypeRepo  services.IToolTypeRepo
}

func NewProductService(
	productRepo services.IProductRepo,
	componentRepo services.IComponentRepo,
	toolTypeRepo services.IToolTypeRepo,
) *productService {
	service := &productService{
		productRepo:   productRepo,
		componentRepo: componentRepo,
		toolTypeRepo:  toolTypeRepo,
	}
	return service
}
