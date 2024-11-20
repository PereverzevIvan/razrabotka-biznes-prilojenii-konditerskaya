import {
  useState,
  createContext,
  useContext,
  useEffect,
  ReactNode,
} from "react";
import { jwtDecode } from "jwt-decode";
import Cookies from "js-cookie";
import { baseURL, routePaths as rp } from "../../configs";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import { TAuthContext, TCredentialsForLogin } from "..";

const CStorageItems = {
  userID: "userID",
  role: "role",
  isAuth: "isAuth",
};

// Контекст для хранения данных авторизации
const AuthContext = createContext<TAuthContext | null>(null);

export function AuthContextProvider({ children }: { children: ReactNode }) {
  const [userID, setUserID] = useState(
    sessionStorage.getItem(CStorageItems.userID)
  );
  const [role, setRole] = useState(sessionStorage.getItem(CStorageItems.role));
  const [isAuth, setIsAuth] = useState(
    sessionStorage.getItem(CStorageItems.isAuth)
  );

  console.log("Auth context is ready", userID, role, isAuth);

  const api = axios.create({
    baseURL: baseURL,
    withCredentials: true,
  });

  const navigate = useNavigate();

  useEffect(() => {
    const accessToken = Cookies.get("access-token");
    if (accessToken) {
      decodeToken(accessToken);
      return;
    }
  }, []);

  // Функция для декодирования токена
  function decodeToken(token: string) {
    type TDecodedToken = {
      id: string;
      role: string;
    };

    try {
      const decoded: TDecodedToken = jwtDecode(token);

      setUserID(decoded.id);
      setRole(decoded.role);
      setIsAuth("true");

      sessionStorage.setItem(CStorageItems.isAuth, "true");
      sessionStorage.setItem(CStorageItems.userID, decoded.id);
      sessionStorage.setItem(CStorageItems.role, decoded.role);
    } catch (error) {
      throw error;
    }
  }

  async function login(credentials: TCredentialsForLogin) {
    try {
      const response = await api.post("/login", credentials); // Получаем токен

      const accessToken = Cookies.get("access_token"); // Достаем токен из кук
      if (typeof accessToken === "undefined") {
        // Проверяем наличие токена
        throw new Error("Не удалось получить токен");
      }

      // Декодируем токен
      decodeToken(accessToken);

      console.log("login успешен");
      return response;
    } catch (error: any) {
      if (error.response) {
        if (error.response.status == 403) {
          console.error(
            "Данные входа корректны. Однако вы заблокированы и не можете войти в систему"
          );
        }
      }
      console.error("Не удалось выполнить login: ", error);

      // Если произошла ошибка, то очищаем данные авторизации
      clearAuth();
      return await Promise.reject(error);
    }
  }

  // Функция для очистки данных авторизации
  function clearAuth() {
    setUserID(null);
    setRole(null);
    setIsAuth(null);
    Cookies.remove("access-token");
    Cookies.remove("refresh-token");
    sessionStorage.removeItem(CStorageItems.isAuth);
    sessionStorage.removeItem(CStorageItems.userID);
    sessionStorage.removeItem(CStorageItems.role);
    console.log("Очистка данных авторизации успешна");
  }

  async function logout() {
    try {
      const response = await api.get("/logout");
      clearAuth();
      console.log("Logout успешен", userID, role);
      return response;
    } catch (error) {
      console.error("Не удалось выполнить logout: ", error);

      // Если произошла ошибка, то очищаем данные авторизации
      clearAuth();
      navigate(rp.login.path);
      return await Promise.reject(error);
    }
  }

  function refresh() {
    return api
      .get("/refresh")
      .then(() => {
        const accessToken = Cookies.get("access-token");
        if (typeof accessToken === "undefined") {
          throw new Error("Токен должен быть строкой");
        }

        decodeToken(accessToken);

        console.log("refresh токенов успешен");
        return true;
      })
      .catch((error) => {
        console.error("Не удалось сделать refresh токенов", error);
        return false;
      });
  }

  api.interceptors.response.use(
    (response: any) => response,
    async (error: any) => {
      const originalRequest = error.config;

      if (
        error.response &&
        error.response.status == 401 &&
        !originalRequest._retry
      ) {
        originalRequest._retry = true; // Это нужно для того, чтобы обработка ошибки не ушла в вечную рекурсию
        const refreshResult = await refresh();
        if (refreshResult) return await api(originalRequest);
      }
      clearAuth();
      return Promise.reject(error);
    }
  );

  return (
    <AuthContext.Provider
      value={{ userID, role, login, logout, refresh, isAuth }}
    >
      {children}
    </AuthContext.Provider>
  );
}

export function useAuthContext() {
  return useContext(AuthContext) as TAuthContext;
}
