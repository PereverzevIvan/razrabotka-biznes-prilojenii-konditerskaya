package repos_mysql_user

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/models"

func (r *UserRepo) GetByID(user_id int) (*models.User, error) {

	var user *models.User
	err := r.db.
		First(&user, user_id).
		Error

	if err != nil {
		return nil, err
	}

	return user, nil
}
