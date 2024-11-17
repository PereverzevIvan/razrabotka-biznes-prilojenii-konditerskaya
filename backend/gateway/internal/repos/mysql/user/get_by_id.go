package repos_mysql_user

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/models"

func (r *userRepo) GetByID(user_id int) (*models.User, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	var user *models.User
	err := r.Conn.
		First(&user, user_id).
		Error

	if err != nil {
		return nil, err
	}

	return user, nil
}
