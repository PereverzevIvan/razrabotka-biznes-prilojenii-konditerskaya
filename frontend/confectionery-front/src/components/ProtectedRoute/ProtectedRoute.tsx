import { Navigate } from "react-router-dom";
import { routePaths as rp } from "../../configs";
import { useAuthContext } from "../../contexts";

type TProtectedRouteProps = {
  children: React.ReactNode;
  requiredRoles: string[];
};

/** Компонент для защиты маршрутов */
export function ProtectedRoute(props: TProtectedRouteProps) {
  const { isAuth, role } = useAuthContext();
  const needAuth = props.requiredRoles.length > 0;

  // Если пользователь не авторизован
  if (needAuth && !isAuth) {
    return <Navigate to={rp.forbiden.path} />;
  }

  // Если список ролей не пустой и текущая роль не входит в список
  if (needAuth && role && !props.requiredRoles.includes(role)) {
    return <Navigate to={rp.forbiden.path} />;
  }

  return props.children;
}
