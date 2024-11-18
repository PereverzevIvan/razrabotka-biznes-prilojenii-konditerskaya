package repos_mysql_purchased_component

func (purchasedComponentRepo *purchasedComponentRepo) CountRemainingComponents(component_id int) (int64, error) {
	var count_left int64
	err := purchasedComponentRepo.conn.
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
