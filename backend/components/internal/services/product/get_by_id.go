package services_product

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (service *productService) GetByID(id int) (*models.Product, error) {
	return service.productRepo.GetByID(id)
}
