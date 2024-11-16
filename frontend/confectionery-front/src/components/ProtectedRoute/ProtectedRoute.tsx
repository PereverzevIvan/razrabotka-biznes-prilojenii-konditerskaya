import { Navigate } from "react-router-dom";
import { roles } from "../../configs";
import { routePaths as rp } from "../../configs";

type TProtectedRouteProps = {
  children: React.ReactNode;
  requiredRoles: string[];
};

/** Компонент для защиты маршрутов */
export function ProtectedRoute(props: TProtectedRouteProps) {
  const isAuth = true;
  const role = roles.director;

  // Если пользователь не авторизован
  if (!isAuth) {
    return <Navigate to={rp.login.path} />;
  }

  // Если список ролей не пустой и текущая роль не входит в список
  if (props.requiredRoles.length && !props.requiredRoles.includes(role)) {
    return <Navigate to={rp.forbiden.path} />;
  }

  return props.children;
}
