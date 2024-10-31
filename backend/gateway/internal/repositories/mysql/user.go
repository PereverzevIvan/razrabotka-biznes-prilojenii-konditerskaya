package mysql_repo

import (
	"sync"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/gateway/models"
	"gorm.io/gorm"
)

type userRepo struct {
	Conn *gorm.DB
	m    *sync.RWMutex
}

func NewUserRepo(conn *gorm.DB) *userRepo {
	return &userRepo{
		Conn: conn,
	}
}

// func (r UserRepo) Create(user *models.User) error {
// 	err := r.Conn.Create(user).Error
// 	if IsUniqueConstraintError(err) {
// 		return models.ErrUnique
// 	}

// 	// if IsForeignKeyConstraintError(err) &&
// 	// 	strings.Contains(err.Error(), "`FK_Users_Offices` FOREIGN KEY (`OfficeID`) foreignKey `offices` (`ID`))") {
// 	// 	return models.ErrFK
// 	// }

// 	return err
// }

// func (r UserRepo) Update(user *models.User) error {

// 	err := r.Conn.Model(&user).
// 		Select("FirstName", "LastName", "Email", "RoleID", "OfficeID").
// 		Updates(*user).Error

// 	if IsUniqueConstraintError(err) {
// 		return models.ErrUnique
// 	}

// 	if IsForeignKeyConstraintError(err) &&
// 		strings.Contains(err.Error(), "`FK_Users_Offices` FOREIGN KEY (`OfficeID`) foreignKey `offices` (`ID`))") {
// 		return models.ErrFKOfficeIDNotFound
// 	}

// 	return err
// }

// func (r UserRepo) UpdateActive(user_id int, is_active bool) error {
// 	err := r.Conn.Model(&models.User{ID: user_id}).
// 		Update("Active", is_active).Error

// 	return err
// }

func (r *userRepo) GetByID(user_id int) (*models.User, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	var user models.User
	err := r.Conn.First(&user, user_id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// func (u UserRepo) GetByEmail(email string) (*models.User, error) {
// 	var user models.User
// 	err := u.Conn.First(&user, "Email = ?", email).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, err
// }

// func (u UserRepo) GetAll(params map[string]string) (*[]models.User, error) {
// 	var users []models.User
// 	// var totalCount int

// 	query := u.Conn.Model(&models.User{})

// 	if params["office_id"] != "" {
// 		query = query.Where("OfficeID = ?", params["office_id"])
// 	}

// 	err := query.Find(&users).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &users, nil
// }

// func (u UserRepo) IsAdmin(user_id int) (bool, error) {
// 	var role_id models.ERole = models.KRoleNone
// 	res := u.Conn.Model(&models.User{}).
// 		Select("RoleID").
// 		Where("id = ?", user_id).
// 		Scan(&role_id)

// 	is_admin := role_id == models.KRoleAdmin
// 	return is_admin, res.Error
// }

// func (u UserRepo) IsActive(user_id int) (bool, error) {
// 	is_active := false
// 	res := u.Conn.Model(&models.User{}).
// 		Select("Active").
// 		Where("id = ?", user_id).
// 		Scan(&is_active)

// 	return is_active, res.Error
// }
