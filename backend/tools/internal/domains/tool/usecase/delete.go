package tool_usecase

func (u *ToolUsecase) Delete(tool_id int) error {
	return u.toolRepo.Delete(tool_id)
}
