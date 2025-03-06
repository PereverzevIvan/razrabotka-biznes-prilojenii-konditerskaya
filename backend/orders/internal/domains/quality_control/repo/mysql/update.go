package quality_control_repo_mysql

import (
	quality_control_params "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/domains/quality_control/params"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"
	"github.com/gofiber/fiber/v3/log"
)

func (repo *QualityControlRepo) Update(params *quality_control_params.UpdateParams) (*models.QualityControl, error) {

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

	err := repo.db.
		Model(&quality_control).
		Updates(update_fields_map).
		Error
	if err != nil {
		return nil, err
	}

	err = repo.db.First(&quality_control).Error
	if err != nil {
		return nil, err
	}

	return quality_control, nil
}
