package repos_mysql_user

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/models"

func (r *UserRepo) GetByLogin(login string) (*models.User, error) {

	var user *models.User
	err := r.db.
		Where("login = ?", login).
		First(&user).
		Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
