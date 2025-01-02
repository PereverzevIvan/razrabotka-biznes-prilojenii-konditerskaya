package repos_mysql_quality_control

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"

func (repo *qualityControlRepo) GetByOrderID(orderID int) ([]models.QualityControl, error) {
	var quality_controls []models.QualityControl

	err := repo.conn.
		Where("order_id = ?", orderID).
		Find(&quality_controls).
		Error
	if err != nil {
		return nil, err
	}

	return quality_controls, nil
}
