import { useNavigate } from "react-router-dom";
import { Button } from "../../../components/Button";
import { routePaths } from "../../../configs";
import { useAuthContext, useToastContext } from "../../../contexts";

export function LogoutButton() {
  const { logout } = useAuthContext();
  const { addToast } = useToastContext();
  const navigate = useNavigate();

  function handleLogout() {
    logout()
      .then(() => {
        addToast("Вы вышли из системы");
        navigate(routePaths.login.path);
      })
      .catch((error) => {
        console.error(`Ошибка выхода из системы ${error}`);
      });
  }

  return (
    <Button color="red" onClick={() => handleLogout()}>
      Выйти
    </Button>
  );
}
