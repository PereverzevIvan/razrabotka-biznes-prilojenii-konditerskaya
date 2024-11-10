package services_purchased_component

func (service *purchasedComponentService) DeleteEmpty(purchased_component_id int) error {
	return service.purchasedComponentRepo.DeleteEmpty(purchased_component_id)
}
