package services_product

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/services"

type productService struct {
	productRepo services.IProductRepo
}

func NewProductService(repo services.IProductRepo) *productService {
	service := &productService{
		productRepo: repo,
	}
	return service
}
