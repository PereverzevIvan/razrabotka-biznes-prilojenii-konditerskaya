package purchased_component_repo_mysql

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	repos_mysql "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/repos/mysql"
)

func (r *PurchasedComponentRepo) Create(purchased_component *models.PurchasedComponent) error {

	err := r.db.
		Create(&purchased_component).
		Error

	if err != nil {
		return err
	}

	if repos_mysql.IsForeignKeyConstraintError(err) {
		return models.ErrFK
	}

	return nil
}
