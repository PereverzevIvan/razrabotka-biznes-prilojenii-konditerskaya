package repos_mysql_user

import "github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/models"

func (r *userRepo) GetByLogin(login string) (*models.User, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	var user *models.User
	err := r.Conn.
		Where("login = ?", login).
		First(&user).
		Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
