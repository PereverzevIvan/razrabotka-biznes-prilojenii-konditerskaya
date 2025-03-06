package product_usecase

func (u *ProductUsecase) CalcComponentsMaxDeliveryTime(component_ids []int) (int, error) {
	var components_max_deliver_time int = 0

	for _, component_id := range component_ids {
		delivery_time_minutes, err := u.supplierComponentRepo.FastestDeliveryComponentTime(component_id)
		if err != nil {
			return 0, err
		}

		components_max_deliver_time = max(delivery_time_minutes, components_max_deliver_time)

	}

	return components_max_deliver_time, nil
}
