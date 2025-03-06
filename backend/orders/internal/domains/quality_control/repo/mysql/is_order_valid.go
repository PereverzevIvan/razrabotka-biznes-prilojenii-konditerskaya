package quality_control_repo_mysql

func (r *QualityControlRepo) IsOrderValid(orderID int) (bool, error) {

	var count_unsatisfied int64

	err := r.db.
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
