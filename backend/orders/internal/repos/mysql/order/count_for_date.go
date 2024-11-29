package repos_mysql_order

import (
	"time"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/orders/internal/models"
)

func (repo *orderRepo) CountForDate(date time.Time) (int64, error) {
	var count int64

	date_str := date.Format("2006-01-02")

	err := repo.conn.
		Table((&models.Order{}).TableName()).
		Where("DATE(created_at) = ?", date_str).
		Count(&count).
		Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
