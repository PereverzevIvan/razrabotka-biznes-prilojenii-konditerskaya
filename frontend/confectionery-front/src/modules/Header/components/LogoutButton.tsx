import { Button } from "../../../components/Button";
import { useAuthContext, useToastContext } from "../../../contexts";

export function LogoutButton() {
  const { logout } = useAuthContext();
  const { addToast } = useToastContext();

  function handleLogout() {
    logout()
      .then(() => {
        addToast("Вы вышли из системы");
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
