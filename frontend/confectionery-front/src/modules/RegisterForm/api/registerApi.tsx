import { TRegisterForm } from "../registerForm.types";
import { apiClient, routePaths } from "../../../configs";

export async function postRegister(postData: TRegisterForm) {
  const { data } = await apiClient.post(routePaths.register.path, postData);
  return data;
}
