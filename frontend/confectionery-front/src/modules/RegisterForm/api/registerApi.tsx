import { TRegisterForm } from "../registerForm.types";
import { apiClient, routePaths } from "../../../configs";

export function register(data: TRegisterForm) {
  return apiClient.post(routePaths.register.path, data);
}
