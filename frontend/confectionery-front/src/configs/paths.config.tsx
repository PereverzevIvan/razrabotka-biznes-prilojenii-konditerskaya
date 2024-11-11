import { ROLES } from "./roles.config";

type TPaths = {
  [key: string]: {
    path: string; // путь
    allowedRoles: string[]; // разрешенные роли (если список пустой, то доступ разрешен всем)
    title: string; // название страницы
  };
};

export const PATHS: TPaths = {
  main: {
    path: "/",
    allowedRoles: [],
    title: "Главная",
  },
  login: {
    path: "/login",
    allowedRoles: [],
    title: "Авторизация",
  },
  register: {
    path: "/register",
    allowedRoles: [],
    title: "Регистрация",
  },
};
