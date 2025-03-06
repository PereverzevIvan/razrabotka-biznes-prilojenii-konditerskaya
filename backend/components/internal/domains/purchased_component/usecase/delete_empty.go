package purchased_component_usecase

func (u *PurchasedComponentUsecase) DeleteEmpty(purchased_component_id int) error {
	return u.purchasedComponentRepo.DeleteEmpty(purchased_component_id)
}
