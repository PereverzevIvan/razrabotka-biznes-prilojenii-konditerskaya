import { useNavigate } from "react-router-dom";
import Button from "../../../components/Button/Button";
import { routePaths as rp } from "../../../configs";

export function LogoutButton() {
  const navigate = useNavigate();
  return (
    <Button color="red" onClick={() => navigate(rp.logout.path)}>
      Выйти
    </Button>
  );
}
