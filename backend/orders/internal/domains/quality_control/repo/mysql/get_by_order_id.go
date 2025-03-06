package quality_control_repo_mysql

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"

func (r *QualityControlRepo) GetByOrderID(orderID int) ([]models.QualityControl, error) {
	var quality_controls []models.QualityControl

	err := r.db.
		Where("order_id = ?", orderID).
		Find(&quality_controls).
		Error
	if err != nil {
		return nil, err
	}

	return quality_controls, nil
}
