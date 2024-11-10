package repos_mysql_purchased_component

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/models"
	repos_mysql "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/components/internal/repos/mysql"
)

func (repo *purchasedComponentRepo) Create(purchased_component *models.PurchasedComponent) error {

	err := repo.conn.
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
