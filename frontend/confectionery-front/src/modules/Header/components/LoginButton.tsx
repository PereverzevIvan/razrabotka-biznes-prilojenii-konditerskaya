import { useNavigate } from "react-router-dom";
import { Button } from "../../../components/Button";
import { routePaths as rp } from "../../../configs";

export function LoginButton() {
  const navigate = useNavigate();
  return (
    <Button color="green" onClick={() => navigate(rp.login.path)}>
      Войти
    </Button>
  );
}
