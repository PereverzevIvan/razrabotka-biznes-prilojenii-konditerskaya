package repos_mysql_user

import (
	repos_mysql "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/internal/repos/mysql"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/models"
)

func (r *UserRepo) Create(user *models.User) error {

	err := r.db.
		Create(user).
		Error

	if repos_mysql.IsUniqueConstraintError(err) {
		return models.ErrUnique
	}

	// if IsForeignKeyConstraintError(err) &&
	// 	strings.Contains(err.Error(), "`FK_Users_Offices` FOREIGN KEY (`OfficeID`) foreignKey `offices` (`ID`))") {
	// 	return models.ErrFK
	// }

	return err
}
