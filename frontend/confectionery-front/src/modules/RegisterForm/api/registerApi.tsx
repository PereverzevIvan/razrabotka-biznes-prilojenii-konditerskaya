import { TRegisterForm } from "../registerForm.types";
import { apiClient } from "../../../configs";

export async function postRegister(postData: TRegisterForm) {
  const { data } = await apiClient.post("/register", postData);
  return data;
}
