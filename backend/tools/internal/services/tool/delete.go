package services_tool

func (service *toolService) Delete(tool_id int) error {
	return service.toolRepo.Delete(tool_id)
}
