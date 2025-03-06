package product_usecase

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"

func (u *ProductUsecase) GetAll() ([]models.Product, error) {
	return u.productRepo.GetAll()
}
