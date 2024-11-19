import { Route, Routes } from "react-router-dom";
import { routePaths as rp } from "../../configs";
import { ProtectedRoute } from "../../components/ProtectedRoute";

/** Компонент для отображения страниц по определенным в конфиге маршрутам */
export function CustomRouter() {
  return (
    <Routes>
      {Object.values(rp).map((route, index) => (
        <Route
          key={index}
          path={route.path}
          element={
            <ProtectedRoute requiredRoles={route.allowedRoles}>
              {route.element}
            </ProtectedRoute>
          }
        />
      ))}
    </Routes>
  );
}