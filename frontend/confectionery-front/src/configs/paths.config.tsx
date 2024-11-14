import DecorationsAccountingPage from "../pages/DecorationsAccountingPage/DecorationsAccountingPage";
import ForbidenPage from "../pages/ForbidenPage/ForbidenPage";
import IngredientsAccountingPage from "../pages/IngredientsAccountingPage/IngredientsAccountingPage";
import LoginPage from "../pages/LoginPage/LoginPage";
import MainPage from "../pages/MainPage/MainPage";
import RegisterPage from "../pages/RegisterPage/RegisterPage";
import ToolsAccountingPage from "../pages/ToolsAccountingPage/ToolsAccountingPage";
import { roles } from "./roles.config";

type TPath = {
  path: string; // путь
  allowedRoles: string[]; // разрешенные роли (если список пустой, то доступ разрешен всем)
  title: string; // название страницы
  element: React.ReactNode | null;
};

type TRoutePaths = {
  [key: string]: TPath;
};

export const baseURL = "http://localhost:3000/api";

export const routePaths: TRoutePaths = {
  main: {
    path: "/",
    allowedRoles: [],
    title: "Главная",
    element: <MainPage />,
  },
  login: {
    path: "/login",
    allowedRoles: [],
    title: "Авторизация",
    element: <LoginPage />,
  },
  logout: {
    path: "/logout",
    allowedRoles: [],
    title: "Выход из аккаунта",
    element: null,
  },
  register: {
    path: "/register",
    allowedRoles: [],
    title: "Регистрация",
    element: <RegisterPage />,
  },
  toolsAccounting: {
    path: "/tools/accouning",
    allowedRoles: [roles.director],
    title: "Учет инструментов",
    element: <ToolsAccountingPage />,
  },
  ingredientsAccounting: {
    path: "/ingredients/accounting",
    allowedRoles: [
      roles.director,
      roles.master,
      roles.client_manager,
      roles.purchase_manager,
    ],
    title: "Учет ингредиентов",
    element: <IngredientsAccountingPage />,
  },
  decorationsAccounting: {
    path: "/decorations/accounting",
    allowedRoles: [
      roles.director,
      roles.master,
      roles.client_manager,
      roles.purchase_manager,
    ],
    title: "Учет декораций",
    element: <DecorationsAccountingPage />,
  },
  forbiden: {
    path: "/forbiden",
    allowedRoles: [],
    title: "403",
    element: <ForbidenPage />,
  },
};
