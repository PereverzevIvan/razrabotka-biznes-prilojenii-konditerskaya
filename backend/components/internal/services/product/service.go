package services_product

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/services"

type productService struct {
	productRepo   services.IProductRepo
	componentRepo services.IComponentRepo
}

func NewProductService(
	productRepo services.IProductRepo,
	componentRepo services.IComponentRepo,
) *productService {
	service := &productService{
		productRepo:   productRepo,
		componentRepo: componentRepo,
	}
	return service
}
