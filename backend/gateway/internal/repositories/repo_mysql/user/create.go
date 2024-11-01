package repo_mysql_user

import (
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/internal/repositories/repo_mysql"
	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/models"
)

func (r *userRepo) Create(user *models.User) error {
	r.m.Lock()
	defer r.m.Unlock()

	err := r.Conn.
		Create(user).
		Error

	if repo_mysql.IsUniqueConstraintError(err) {
		return models.ErrUnique
	}

	// if IsForeignKeyConstraintError(err) &&
	// 	strings.Contains(err.Error(), "`FK_Users_Offices` FOREIGN KEY (`OfficeID`) foreignKey `offices` (`ID`))") {
	// 	return models.ErrFK
	// }

	return err
}
