import { useNavigate } from "react-router-dom";
import { Button } from "../../../components/Button";
import { routePaths as rp } from "../../../configs";

export function RegisterButton() {
  const navigate = useNavigate();
  return (
    <Button color="blue" onClick={() => navigate(rp.register.path)}>
      Регистрация
    </Button>
  );
}
