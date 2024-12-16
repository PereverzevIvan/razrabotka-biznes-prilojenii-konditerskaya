import { TRegisterForm } from "../modules/RegisterForm";
import { apiClient } from "./apiClient";

export async function postRegister(postData: TRegisterForm) {
  const { data } = await apiClient.post("/register", postData);
  return data;
}
