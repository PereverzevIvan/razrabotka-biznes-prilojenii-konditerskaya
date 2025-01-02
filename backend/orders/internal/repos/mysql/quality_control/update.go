package repos_mysql_quality_control

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"
	params_quality_control "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models/params/quality_control"
	"github.com/gofiber/fiber/v3/log"
)

func (repo *qualityControlRepo) Update(params *params_quality_control.UpdateParams) (*models.QualityControl, error) {

	update_fields_map := map[string]interface{}{}

	if params.IsSatisfies != nil {
		update_fields_map["is_satisfies"] = params.IsSatisfies
	}

	log.Info(params)

	if params.Comment != "" {
		update_fields_map["comment"] = params.Comment
	}

	if params.Parameter != nil {
		update_fields_map["parameter"] = params.Parameter
	}

	quality_control := &models.QualityControl{
		ID: params.ID,
	}

	err := repo.conn.
		Model(&quality_control).
		Updates(update_fields_map).
		Error
	if err != nil {
		return nil, err
	}

	err = repo.conn.First(&quality_control).Error
	if err != nil {
		return nil, err
	}

	return quality_control, nil
}
