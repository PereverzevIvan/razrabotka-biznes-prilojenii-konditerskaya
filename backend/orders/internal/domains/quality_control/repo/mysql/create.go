package quality_control_repo_mysql

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"

func (r *QualityControlRepo) Create(quality_control *models.QualityControl) (*models.QualityControl, error) {

	err := r.db.Create(quality_control).Error
	if err != nil {
		return nil, err
	}

	return quality_control, nil
}
