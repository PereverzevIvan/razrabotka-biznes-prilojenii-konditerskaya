package purchased_component_repo_mysql

func (r *PurchasedComponentRepo) CountRemainingComponents(component_id int) (int64, error) {
	var count_left int64
	err := r.db.
		Table("purchased_components").
		Select("ifnull(sum(quantity), 0) as count_left").
		Where("component_id = ?", component_id).
		Where("quantity > 0").
		Scan(&count_left).
		Error
	if err != nil {
		return 0, err
	}

	return count_left, nil
}
