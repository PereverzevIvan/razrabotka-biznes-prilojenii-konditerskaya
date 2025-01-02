package repos_mysql_quality_control

func (repo *qualityControlRepo) IsOrderValid(orderID int) (bool, error) {

	var count_unsatisfied int64

	err := repo.conn.
		Table("quality_controls").
		Where("order_id = ?", orderID).
		Where("is_satisfies = ?", false).
		Count(&count_unsatisfied).
		Error
	if err != nil {
		return false, err
	}

	return count_unsatisfied == 0, nil
}
