package repos_mysql_quality_control

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"

func (repo *qualityControlRepo) Create(quality_control *models.QualityControl) (*models.QualityControl, error) {
	
    
    err := repo.conn.Create(quality_control).Error
	if err != nil {
		return nil, err
	}

	return quality_control, nil
}
